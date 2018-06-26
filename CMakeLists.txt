PROJECT(POSEPROTO)

CMAKE_MINIMUM_REQUIRED(VERSION 3.1)

FIND_PACKAGE(Protobuf REQUIRED)
set(SRC
  pose.proto
)
PROTOBUF_GENERATE_CPP(POSE_PROTO_SRC POSE_PROTO_HDR ${SRC})
set(POSE_PROTO_SRC "${POSE_PROTO_SRC}" CACHE INTERNAL "Generated source file list for pose proto")
set(POSE_PROTO_HDR "${POSE_PROTO_HDR}" CACHE INTERNAL "Generated header file list for pose proto")
add_library(poseproto SHARED
  ${POSE_PROTO_SRC}
  ${POSE_PROTO_HDR}
)

target_include_directories(poseproto PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}>
  $<INSTALL_INTERFACE:include/poseproto>
  ${Protobuf_INCLUDE_DIRS}
)

target_link_libraries(poseproto PUBLIC
  ${Protobuf_LIBRARIES}
)

install(TARGETS poseproto EXPORT poseproto
  LIBRARY DESTINATION lib
)
install(DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/ DESTINATION include/poseproto FILES_MATCHING PATTERN "*.h*")
install(EXPORT poseproto DESTINATION share/poseproto/cmake FILE poseprotoConfig.cmake)
