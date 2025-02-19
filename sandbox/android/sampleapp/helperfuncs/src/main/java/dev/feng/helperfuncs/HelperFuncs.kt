package dev.feng.helperfuncs
import dev.feng.helperfuncs.commontypes.Err



import java.net.HttpURLConnection
import java.net.URI
import com.google.gson.Gson


class HelperFuncs {

    companion object{
        fun getResourceAsText(url: String): Pair<String, Err?> {
            val connection = URI(url).toURL().openConnection() as HttpURLConnection
            return try {
                connection.requestMethod = "GET"
                connection.connect()
                val text = connection.inputStream.bufferedReader().use { it.readText() }
                Pair(text, null)
            } catch (e: Exception) {
                e.printStackTrace()
                Pair("", object : Err {
                    override fun error(): String = e.message ?: "Error fetching resource"
                })
            } finally {
                connection.disconnect()
            }
        }

        fun getResourceAsBytes(url: String): ByteArray {
            return object {}.javaClass.getResource(url).readBytes()
        }

        inline fun <reified T> getResourceWithJson(url: String,json: String): Pair<T?,Err?> {
            val connection = URI(url).toURL().openConnection() as HttpURLConnection
            return try {
                connection.requestMethod = "GET"
                connection.setRequestProperty("Content-Type", "application/json")
                connection.doOutput = true
                connection.outputStream.bufferedWriter().use { it.write(json) }
                connection.connect()
                val text = connection.inputStream.bufferedReader().use { it.readText() }
                val obj: T = Gson().fromJson(text, T::class.java)
                Pair(obj, null)
            } catch (e: Exception) {
                e.printStackTrace()
                Pair(null, object : Err {
                    override fun error(): String = e.message ?: "Error fetching resource"
                })
            } finally {
                connection.disconnect()
            }
        }

        inline fun <reified T> getResource(url: String): Pair<T?,Err?> {
            val connection = URI(url).toURL().openConnection() as HttpURLConnection
            return try {
                connection.requestMethod = "GET"
                connection.connect()
                val text = connection.inputStream.bufferedReader().use { it.readText() }
                val obj: T = Gson().fromJson(text, T::class.java)
                Pair(obj, null)
            } catch (e: Exception) {
                e.printStackTrace()
                Pair(null, object : Err {
                    override fun error(): String = e.message ?: "Error fetching resource"
                })
            } finally {
                connection.disconnect()
            }
        }

        fun <T> toJson(obj: T): String {
            return Gson().toJson(obj)
        }

        inline fun <reified T> fromJson(json: String): T {
            return Gson().fromJson(json, T::class.java)
        }


    }


}