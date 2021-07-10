package framework

//pack len(u16) + cmd(u16)
const PACKAGE_HEADER_LENGTH = 4

//head + body <= 65535
const MAX_PACK_LENGTH = 65535
