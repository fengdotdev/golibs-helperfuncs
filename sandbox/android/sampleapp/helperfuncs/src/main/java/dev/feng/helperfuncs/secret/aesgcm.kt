package dev.feng.helperfuncs.secret


import java.security.SecureRandom
import dev.feng.helperfuncs.commontypes.Err

fun generateKey():ByteArray {
    val key = ByteArray(32)
    SecureRandom().nextBytes(key)
    return key
}


fun generateIV():ByteArray {
    val iv = ByteArray(12)
    SecureRandom().nextBytes(iv)
    return iv
}


fun encodeAESGCM(key: ByteArray, iv: ByteArray, plaintext: ByteArray,additionalData: ByteArray?):Pair<ByteArray,Err?> {
    val cipher = javax.crypto.Cipher.getInstance("AES/GCM/NoPadding")
    val secretKey = javax.crypto.spec.SecretKeySpec(key, "AES")
    val gcmParameterSpec = javax.crypto.spec.GCMParameterSpec(128, iv)
    cipher.init(javax.crypto.Cipher.ENCRYPT_MODE, secretKey, gcmParameterSpec)
    if (additionalData != null){
        cipher.updateAAD(additionalData)
    }
    return Pair(cipher.doFinal(plaintext),null)
}


fun decodeAESGCM(key: ByteArray, iv: ByteArray, cyphertext: ByteArray,additionalData: ByteArray?):Pair<ByteArray,Err?> {

    val cipher = javax.crypto.Cipher.getInstance("AES/GCM/NoPadding")
    val secretKey = javax.crypto.spec.SecretKeySpec(key, "AES")
    val gcmParameterSpec = javax.crypto.spec.GCMParameterSpec(128, iv)
    cipher.init(javax.crypto.Cipher.DECRYPT_MODE, secretKey, gcmParameterSpec)
    if (additionalData != null){
        cipher.updateAAD(additionalData)
    }
    return Pair(cipher.doFinal(cyphertext),null)
}