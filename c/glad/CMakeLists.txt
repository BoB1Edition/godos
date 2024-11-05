cmake_minimum_required(VERSION 3.9)
project(myclib VERSION 1.0.1 DESCRIPTION "myclib for go bin")
include_directories(SYSTEM include)
set(GLAD_LIB src/gl.c)
add_library(glad SHARED ${GLAD_LIB})
set_target_properties(glad PROPERTIES PUBLIC_HEADER "include/glad/gl.h;include/KHR/khrplatform.h")
set_target_properties(
    glad
    PROPERTIES 
    INSTALL_RPATH 
    "${CMAKE_INSTALL_RPATH};${CMAKE_INSTALL_LIBDIR}/glad;${CMAKE_SOURCE_DIR}")  
include(GNUInstallDirs)
configure_file(glad.pc.in glad.pc @ONLY)
install(TARGETS glad
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}/glad
    PUBLIC_HEADER DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/glad)
install(FILES ${CMAKE_BINARY_DIR}/include/KHR/khrplatform.h DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/KHR)
install(FILES ${CMAKE_BINARY_DIR}/glad.pc DESTINATION ${CMAKE_INSTALL_DATAROOTDIR}/pkgconfig)