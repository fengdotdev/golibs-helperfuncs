package dev.feng.helperfuncs
import dev.feng.helperfuncs.commontypes.Err


import kotlinx.serialization.encodeToString
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.json.Json

class HelperFuncs {

    companion object{
        fun getResourceAsText(url: String): String {
            return object {}.javaClass.getResource(url).readText()
        }

        fun getResourceAsBytes(url: String): ByteArray {
            return object {}.javaClass.getResource(url).readBytes()
        }

        fun <T> getResource(url: String): Pair<T?,Err?> {
            try {
                val data = object {}.javaClass.getResource(url).readBytes()
                return Pair(data as T,null)
            } catch (e: Exception) {
                return Pair( null, object : Err {
                    override fun error(): String {
                        return e.message.toString()
                    }
                })

            }
        }

        inline fun <reified T> toJson(obj: T): String {
            return Json.encodeToString(obj)
        }

        inline fun <reified T> fromJson(json: String): T {
            return Json.decodeFromString(json)
        }
    }


}