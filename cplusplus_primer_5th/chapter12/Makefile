.PHONY:all
ALL_BINARYS := test_array\
				cpp_12_23\
				cpp_12_24\
				cpp_12_26
all : $(ALL_BINARYS)

CXXFLAGS = -O0 -g -Wall -Werror -std=c++11

%:
	g++ $(CXXFLAGS) -o $@ $(addsuffix .cpp,$@)

# Use the following command to disassemble executable code
#@objdump -d $(BINARY) > test_array_dump
clean:
	rm -rf $(ALL_BINARYS)
	rm -rf test_array_dump
