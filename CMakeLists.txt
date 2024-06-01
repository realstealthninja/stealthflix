cmake_minimum_required(VERSION 3.28.10)


project(stealthflix
  LANGUAGES CXX
  VERSION 0.1.0
  DESCRIPTION "A small website for all your streamig needs"
)

# Cmake settings
set(CMAKE_EXPORT_COMPILE_COMMANDS TRUE)


# add sub modules
set(CROW_FEATURES  ssl  compression)
add_subdirectory("extern/Crow/")

## MEGA.nz SDK
add_subdirectory("extern/sdk/")

# adding executable
add_executable(stealthflix main.cpp)

# link packages 
target_link_libraries(stealthflix Crow::Crow MEGA::SDKlib)


# adding project sources and headers


# copying over directories
