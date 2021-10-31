package mn.haruulzangi.composeyourflag

import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import mn.haruulzangi.composeyourflag.core.pbkdf2Calculate
import mn.haruulzangi.composeyourflag.ui.screens.PasswordScreen
import mn.haruulzangi.composeyourflag.ui.screens.PasswordViewModel
import mn.haruulzangi.composeyourflag.ui.theme.ComposeYourFlagTheme
import mn.haruulzangi.composeyourflag.utils.bytesToHex
import mn.haruulzangi.composeyourflag.utils.hexToByteArray
import java.lang.Exception
import java.security.MessageDigest
import java.security.SecureRandom
import javax.crypto.Cipher
import javax.crypto.KeyGenerator
import javax.crypto.SecretKey
import javax.crypto.spec.IvParameterSpec
import javax.crypto.spec.SecretKeySpec

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        setContent {
            ComposeYourFlagTheme {
                val passwordViewModel = viewModel<PasswordViewModel>()
                PasswordScreen(
                    state = passwordViewModel.state.collectAsState().value,
                    onSubmitPassword = passwordViewModel::submit,
                )
            }
        }
    }
}
