package dev.feng.helperfuncs.data

import dev.feng.helperfuncs.commontypes.Err


fun getSHA256(data: String): String {
    val bytes = data.toByteArray()
    val md = java.security.MessageDigest.getInstance("SHA-256")
    val digest = md.digest(bytes)
    return digest.fold("", { str, it -> str + "%02x".format(it) })
}


fun getSHA256Bytes(data: ByteArray): String {
    val md = java.security.MessageDigest.getInstance("SHA-256")
    val digest = md.digest(data)
    return digest.fold("", { str, it -> str + "%02x".format(it) })
}


fun validateSHA256(data: String, hash: String): Err? {
    val hashData = getSHA256(data)
    if (hashData == hash) {
        return null
    }
    return object : Err {
        override fun error(): String {
            return "Hashes do not match"
        }
    }
}


fun validateSHA256Bytes(data: ByteArray, hash: String): Err? {
    val hashData = getSHA256Bytes(data)
    if (hashData == hash) {
        return null
    }
    return object : Err {
        override fun error(): String {
            return "Hashes do not match"
        }
    }
}


fun sHA256isValid(data: String, hash: String): Boolean {
    val hashData = getSHA256(data)
    return hashData == hash
}