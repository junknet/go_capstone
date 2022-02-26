#----------------------------------------------------------------
# Generated CMake target import file.
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "capstone::capstone-static" for configuration ""
set_property(TARGET capstone::capstone-static APPEND PROPERTY IMPORTED_CONFIGURATIONS NOCONFIG)
set_target_properties(capstone::capstone-static PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_NOCONFIG "C"
  IMPORTED_LOCATION_NOCONFIG "${_IMPORT_PREFIX}/lib/libcapstone.a"
  )

list(APPEND _IMPORT_CHECK_TARGETS capstone::capstone-static )
list(APPEND _IMPORT_CHECK_FILES_FOR_capstone::capstone-static "${_IMPORT_PREFIX}/lib/libcapstone.a" )

# Import target "capstone::capstone-shared" for configuration ""
set_property(TARGET capstone::capstone-shared APPEND PROPERTY IMPORTED_CONFIGURATIONS NOCONFIG)
set_target_properties(capstone::capstone-shared PROPERTIES
  IMPORTED_LOCATION_NOCONFIG "${_IMPORT_PREFIX}/lib/libcapstone.so.5.0.0"
  IMPORTED_SONAME_NOCONFIG "libcapstone.so.5"
  )

list(APPEND _IMPORT_CHECK_TARGETS capstone::capstone-shared )
list(APPEND _IMPORT_CHECK_FILES_FOR_capstone::capstone-shared "${_IMPORT_PREFIX}/lib/libcapstone.so.5.0.0" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)
