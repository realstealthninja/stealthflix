cmake_minimum_required(VERSION 3.28.10)


project(stealthflix
  LANGUAGES CXX
  VERSION 0.1.0
  DESCRIPTION "A small website for all your streamig needs"
)

# add sub modules
set(CROW_FEATURES  ssl  compression)
add_subdirectory("extern/Crow/")

# adding executable
add_executable(stealthflix main.cpp)

# link packages 
target_link_libraries(stealthflix Crow::Crow)


# adding project sources and headers


