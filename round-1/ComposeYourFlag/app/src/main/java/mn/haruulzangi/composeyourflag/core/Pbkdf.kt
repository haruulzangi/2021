package mn.haruulzangi.composeyourflag.core

import mn.haruulzangi.composeyourflag.utils.bytesToHex
import mn.haruulzangi.composeyourflag.utils.hexToByteArray
import java.security.SecureRandom
import javax.crypto.SecretKeyFactory
import javax.crypto.spec.PBEKeySpec

private fun pbkdf2(salt: ByteArray, data: String) =
    SecretKeyFactory
        .getInstance("PBKDF2WithHmacSHA1")
        .generateSecret(
            PBEKeySpec(data.toCharArray(), salt, 1000, 256)
        )
        .encoded

fun pbkdf2Calculate(data: String): String {
    val salt = ByteArray(8).apply {
        SecureRandom.getInstance("SHA1PRNG").nextBytes(this)
    }
    return "${bytesToHex(salt)}:${bytesToHex(pbkdf2(salt, data))}"
}
fun pbkdf2Verify(password: String, digest: String): Boolean =
    try {
        val salt = hexToByteArray(digest.split(':')[0])
        val hash = hexToByteArray(digest.split(':')[1])
        pbkdf2(salt, password).contentEquals(hash)
    } catch (e: Exception) {
        false
    }
