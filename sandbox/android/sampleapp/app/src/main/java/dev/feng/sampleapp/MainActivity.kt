package dev.feng.sampleapp

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.lifecycle.lifecycleScope
import dev.feng.sampleapp.ui.theme.SampleappTheme
import dev.feng.helperfuncs.HelperFuncs
import dev.feng.helperfuncs.commontypes.Err
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext


class MainActivity : ComponentActivity() {

    private fun launchCoroutine(callback: suspend CoroutineScope.() -> Unit) {
        lifecycleScope.launch {
            callback()
        }
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            SampleappTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    val msg = remember { mutableStateOf(helper()) }

                    LaunchedEffect(Unit) {
                        msg.value = helperCoroutine()
                    }
                    Column(
                        modifier = Modifier
                            .fillMaxSize()
                            .padding(innerPadding),
                        verticalArrangement = Arrangement.Center,
                        horizontalAlignment = Alignment.CenterHorizontally
                    ) {
                        Greeting(msg.value, Modifier.padding(16.dp))
                        UpdateButton {
                            msg.value = "updating..."
                            launchCoroutine {
                                val result = helperCoroutine()
                                msg.value = result
                            }
                        }
                    }
                }
            }
        }
    }
}

@Composable
fun UpdateButton(callback: () -> Unit) {
    Button(onClick = {
        callback()
    }) {
        Text("Update")
    }
}

@Composable
fun Greeting(msg: String, modifier: Modifier) {
    Text(
        text = msg,
        modifier = modifier
    )
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    SampleappTheme {
        Greeting(
            msg = "Hello World!",
            modifier = Modifier.padding(16.dp)
        )
    }
}

suspend fun helperCoroutine( ): String = withContext(Dispatchers.IO) {
    helper()
}

fun helper(): String {

    val ip = "192.168.100.76:8080"
    val endpoint = "http://$ip/ping"
    val (result: String, err: Err?) = HelperFuncs.getResourceAsText(endpoint)
    if (err != null) {
        return err.error()
    }
    return result
}

