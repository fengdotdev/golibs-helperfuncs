package dev.feng.helperfuncs.data

import java.util.Base64


// Encode data to base64 and return the encoded string
 fun encode64(data: String): String {
     if (data.isEmpty()) {
         return ""
     }
    return Base64.getEncoder().encodeToString(data.toByteArray())
}


// Encode data to base64 and return the encoded string
fun encode64Bytes(data: ByteArray): String {
    return Base64.getEncoder().encodeToString(data)
}


// Decode base64 data to string and return a pair with the result and the decoded string, if the decoding was successful (true) or not (false)
fun decode64(data: String): Pair<Boolean, String> {
    return try {
        val decodedBytes = Base64.getDecoder().decode(data)
        Pair(true, String(decodedBytes))
    } catch (e: Exception) {
        Pair(false, "")
    }
}

// Decode base64 data to bytes array and return a pair with the result and the decoded bytes array, if the decoding was successful (true) or not (false)
fun decode64Bytes(data: String): Pair<Boolean, ByteArray> {
    return try {
        val decodedBytes = Base64.getDecoder().decode(data)
        Pair(true, decodedBytes)
    } catch (e: Exception) {
        Pair(false, ByteArray(0))
    }
}