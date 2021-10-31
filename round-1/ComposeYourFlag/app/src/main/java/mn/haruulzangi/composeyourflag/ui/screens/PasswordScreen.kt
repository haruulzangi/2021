package mn.haruulzangi.composeyourflag.ui.screens

import androidx.compose.animation.AnimatedVisibility
import androidx.compose.animation.ExperimentalAnimationApi
import androidx.compose.animation.slideInVertically
import androidx.compose.animation.slideOutVertically
import androidx.compose.foundation.layout.*
import androidx.compose.material.*
import androidx.compose.runtime.*
import androidx.compose.ui.ExperimentalComposeUiApi
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalSoftwareKeyboardController
import androidx.compose.ui.platform.SoftwareKeyboardController
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import mn.haruulzangi.composeyourflag.ui.theme.ComposeYourFlagTheme

@OptIn(ExperimentalComposeUiApi::class, ExperimentalAnimationApi::class)
@Composable
fun PasswordScreen(state: PasswordState, onSubmitPassword: (String) -> Unit) {
    Scaffold { padding ->
        Column(verticalArrangement = Arrangement.SpaceBetween, modifier = Modifier.fillMaxSize()) {
            Column(
                Modifier
                    .padding(padding)
                    .padding(horizontal = 20.dp)
            ) {
                Spacer(Modifier.height(10.dp))
                Text(text = "Please enter password")
                var password by remember { mutableStateOf("") }
                OutlinedTextField(
                    value = password,
                    singleLine = true,
                    onValueChange = { password = it },
                )
                Spacer(Modifier.height(10.dp))

                val keyboardController = LocalSoftwareKeyboardController.current
                Button(onClick = {
                    keyboardController?.hide()
                    onSubmitPassword(password)
                }) {
                    Text(text = "Check")
                }
            }
            Column {
                AnimatedVisibility(
                    visible = state is PasswordState.Invalid,
                    enter = slideInVertically(initialOffsetY = { it }),
                    exit = slideOutVertically(targetOffsetY = { it }),
                ) {
                    Snackbar(modifier = Modifier.padding(16.dp)) {
                        Text("Invalid password")
                    }
                }
                AnimatedVisibility(
                    visible = state is PasswordState.Success,
                    enter = slideInVertically(initialOffsetY = { it }),
                    exit = slideOutVertically(targetOffsetY = { it }),
                ) {
                    Snackbar(modifier = Modifier.padding(16.dp)) {
                        Text("Congratz! Flag is: ${(state as PasswordState.Success).flag}")
                    }
                }
            }
        }
    }
}

@Preview
@Composable
fun PasswordScreenPreview() {
    ComposeYourFlagTheme {
        PasswordScreen(state = PasswordState.Idle, onSubmitPassword = {})
    }
}