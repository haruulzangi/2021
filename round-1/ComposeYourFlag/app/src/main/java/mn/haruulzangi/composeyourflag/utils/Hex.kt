package mn.haruulzangi.composeyourflag.utils

fun bytesToHex(bytes: ByteArray) = bytes.joinToString("") { "%02x".format(it) }

fun hexToByteArray(hex: String): ByteArray {
    require(hex.length % 2 == 0) { "String must have an even length" }
    return hex.chunked(2).map { it.toInt(16).toByte() }.toByteArray()
}

fun String.toHex(): String = this.toByteArray(Charsets.US_ASCII).let(::bytesToHex)