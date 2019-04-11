#!/usr/bin/env python
import numpy as np
import matplotlib.pyplot as plt
import csv
import re
import pprint

# with open('bench.log.csv', newline='') as csvfile:
#     reader = csv.DictReader(csvfile, delimiter=',')
#     for row in reader:
#         print(row)

csvfile = open('bench.log.csv', newline='')
reader = csv.DictReader(csvfile, delimiter=',')
# for row in reader:
#     print(row["Test"].strip())

tests = {}
# Group results by test name:
for row in reader:
    name = row["Test"].strip()
    name = name.replace("Benchmark", "")

    runtime = float(row["ns/op"].strip())
    if name not in tests:
        tests[name] = {
            "runtimes": []
        }
    tests[name]["runtimes"].append(runtime)

print(",".join(["test", "mean", "avg", "median", "std",
                "variance", "pcnt_of_mean", "pcnt_of_avg"]))
for test in tests:
    v = tests[test]
    v["runtimes"] = np.true_divide(v["runtimes"], 1000000)  # convers ns to ms
    v["mean"] = np.mean(v["runtimes"])
    v["avg"] = np.average(v["runtimes"])
    v["median"] = np.median(v["runtimes"])
    v["std"] = np.std(v["runtimes"])  # std = sqrt(mean(abs(x - x.mean())**2))
    v["variance"] = np.var(v["runtimes"])  # v["std"]**2
    v["pcnt_of_mean"] = v["std"]/v["mean"]*100
    v["pcnt_of_avg"] = v["std"]/v["avg"]*100

    # fields = ["test", "mean", "avg", "median", "std", "variance", "pcnt_of_mean", "pcnt_of_avg"]
    # "{0:.4f}".format(x)
    vals = [x if isinstance(x, str) else "{}".format(x) for x in [test, v["mean"], v["avg"], v["median"], v["std"],
                                                                  v["variance"], v["pcnt_of_mean"], v["pcnt_of_avg"]]]
    # for field, val in zip(fields, vals):
    #     print()
    stmt = ",".join(vals)
    print(stmt)

    # print(test)
    # pprint.pprint(v)

# for test in tests:
#     v = tests[test]
#     plt.figure()
#     plt.errorbar()


def doPlot(testName, protoData, jsonData):
    protoY = [tests[test]["mean"] for test in tests if test in protoData]
    protoYerr = [tests[test]["std"] for test in tests if test in protoData]
    # protoX = [test for test in tests if test in protoData]

    jsonY = [tests[test]["mean"] for test in jsonData if test in jsonData]
    jsonYerr = [tests[test]["std"] for test in jsonData if test in jsonData]
    # jsonX = [test for test in jsonData if test in jsonData]

    # get only the numeric values
    x = [(re.search(r'\d+', test).group()) for test in protoData]

    #     x = ["1" + test.split("1")[1].replace("-8", "") for test in protoData]

    fig, ax = plt.subplots()
    ax.errorbar(x, protoY, yerr=protoYerr, fmt="-o",
                capsize=4, label="proto")
    #ax.plot(x, np.poly1d(np.polyfit(x, protoY, 2))(x), "--")

    ax.errorbar(x, jsonY, yerr=jsonYerr, fmt="-*",
                capsize=4, label="JSON")
    #ax.plot(x, np.poly1d(np.polyfit(x, jsonY, 2))(x), "--")

    ax.set_title(testName)
    ax.yaxis.grid(True)

    ax.legend()

    plt.xlabel("object size")
    plt.ylabel("ms/op")

    plt.savefig(testName.replace(" ", "-")+'-plot.png')
    plt.show()

    #distance = np.subtract(jsonY, protoY)
    pcnt_distance = np.multiply(np.true_divide(protoY, jsonY), 100)
    y_pos = np.arange(len(x))
    plt.bar(y_pos, pcnt_distance, align='center')
    plt.xticks(y_pos, x)
    plt.ylabel('procent (%)')
    plt.xlabel("object size")
    plt.title('Mean forskel i procent')
    plt.savefig(testName.replace(" ", "-")+'-diff-plot.png')
    plt.show()
    #fig2, ax2 = plt.subplots()
    #ax.bar(y_pos, distance)
    # ax.ylabel("Distance")
    # ax.set_title("")


protos = [test for test in tests if "Proto" in test]
jsons = [test for test in tests if "JSON" in test]


protoUnmarshal = []
jsonUnmarshal = []

protoMarshal = []
jsonMarshal = []

for proto, json in zip(protos, jsons):
    name = proto.replace("Proto", " ").replace(
        "-8", "").replace("1", " 1").replace("5", " 5")

    if "Simple" in name:
        continue  # we dont want to group the simple benchmark

    if "UnMarshal" in name:
        # it is an unmarshal test
        protoUnmarshal.append(proto)
        jsonUnmarshal.append(json)
    else:
        protoMarshal.append(proto)
        jsonMarshal.append(json)
    # doPlot(name, protoTests, jsonTests)

doPlot("Unmarshal Nested", protoUnmarshal, jsonUnmarshal)
doPlot("Marshal Nested", protoMarshal, jsonMarshal)
