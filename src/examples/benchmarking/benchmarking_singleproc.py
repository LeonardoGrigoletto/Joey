import pandas as pd
import time, os

if __name__ == "__main__":
    dir = os.path.dirname(os.path.abspath(__file__))
    dataframe = pd.read_csv(os.path.join(dir, "benchmarking_data.csv"))

    print("----- Benchmarking Adding Columns -----")
    start_time = time.time()
    dataframe["charge"] = dataframe["charge"] + dataframe["walltime"]
    print(f"{(time.time() - start_time) * 1000} ms")

    print("----- Benchmarking Subtracting Columns -----")
    start_time = time.time()
    dataframe["charge"] = dataframe["charge"] - dataframe["walltime"]
    print(f"{(time.time() - start_time) * 1000} ms")

    print("----- Benchmarking Multiplying Columns -----")
    start_time = time.time()
    dataframe["charge"] = dataframe["charge"] * dataframe["walltime"]
    print(f"{(time.time() - start_time) * 1000} ms")



