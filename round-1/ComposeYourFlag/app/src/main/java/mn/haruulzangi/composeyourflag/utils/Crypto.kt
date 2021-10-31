package mn.haruulzangi.composeyourflag.utils

import javax.crypto.SecretKey
import javax.crypto.spec.SecretKeySpec

fun ByteArray.asSecretKey(algorithm: String = "AES"): SecretKey =
    SecretKeySpec(this, 0, size, algorithm)