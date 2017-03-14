BINARY = test_array
CXXFLAGS = -O0 -g -Wall -Werror -std=c++11

$(BINARY):
	g++ $(CXXFLAGS) -o $@ $(filter %.cpp,$^)
	@objdump -d $(BINARY) > test_array_dump
test_array: test_array.cpp

clean:
	rm -rf $(BINARY)
	rm -rf test_array_dump
