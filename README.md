# Generating Deadlock-Free and Live Go Code from Unbounded Multiparty Session Protocols

This file contains the tool GoScr, as well as the benchmarks used in the paper
"Generating Deadlock-Free and Live Go Code from Unbounded Multiparty Session
Protocols". 

The benchmark base cases are under 'benchmarks/<n>.<benchmark_name>/base' and
the GoScr versions are under 'benchmarks/<n>.<benchmark_name>/goscr'. To
generate the code and compile the benchmarks, please use the scripts in
'benchmarks/scripts'. The directory 'nuscr' contains the code for our tool.
