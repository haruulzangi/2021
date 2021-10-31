package mn.haruulzangi.composeyourflag.ui.screens

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import mn.haruulzangi.composeyourflag.core.pbkdf2Verify
import mn.haruulzangi.composeyourflag.utils.asSecretKey
import mn.haruulzangi.composeyourflag.utils.hexToByteArray
import java.lang.Exception
import java.security.MessageDigest
import javax.crypto.Cipher
import javax.crypto.SecretKey
import javax.crypto.spec.IvParameterSpec
import javax.crypto.spec.SecretKeySpec

sealed class PasswordState {
    object Idle : PasswordState()
    object Invalid : PasswordState()
    data class Success(val flag: String) : PasswordState()
}

class PasswordViewModel : ViewModel() {
    private val _state = MutableStateFlow<PasswordState>(PasswordState.Idle)
    val state = _state.asStateFlow()

    private fun hash(data: String): ByteArray =
        MessageDigest.getInstance("SHA-256").digest(data.toByteArray())

    private fun decrypt(cipherText: ByteArray, key: SecretKey, IV: ByteArray): String? {
        try {
            val cipher = Cipher.getInstance("AES/CBC/PKCS5Padding")
            val keySpec = SecretKeySpec(key.encoded, "AES")
            val ivSpec = IvParameterSpec(IV)
            cipher.init(Cipher.DECRYPT_MODE, keySpec, ivSpec)
            val decryptedText = cipher.doFinal(cipherText)
            return String(decryptedText)
        } catch (e: Exception) {
            e.printStackTrace()
        }
        return null
    }

    fun submit(password: String) {
        if (state.value is PasswordState.Success) {
            return
        }
        if (state.value is PasswordState.Invalid) {
            _state.value = PasswordState.Idle
        }

        if (!Regex("^\\d{5}$").matches(password) || pbkdf2Verify(password, "")) {
            _state.value = PasswordState.Invalid
            viewModelScope.launch {
                delay(3000)
                _state.value = PasswordState.Idle
            }
            return
        }

        if (!pbkdf2Verify(password, "71ec232b36a74564:4bdea2f4edd6100523e09510d688088470e29cb029e87f3b4d0542207429dec9")) {
            _state.value = PasswordState.Invalid
            viewModelScope.launch {
                delay(1000)
                _state.value = PasswordState.Idle
            }
            return
        }

        val encryptedFlag = hexToByteArray("fc8ff8aa18b06535fff4889c5943ae1e1827b208657e2d9525c3fce6aaa942f0aa5923df1e0caf9ab68adc02338c2dbb")
        val secretKey: SecretKey = hash(password).take(16).toByteArray().asSecretKey()
        val iv = hexToByteArray("9465ed0e693017a5d54ed2346616d6ac")

        val flag = decrypt(encryptedFlag, secretKey, iv).toString()
        _state.value = PasswordState.Success(flag)
    }
}