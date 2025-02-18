package dev.feng.helperfuncs.data

import org.junit.Assert.*
import org.junit.Test

class Base64HelperTest {

    @Test
    fun TestEncodeDecodeWithKnownValues() {
        val originalKnown = "foo"
        val encodedknown = "Zm9v"
        val encoded = encode64(originalKnown)
        assertEquals(encodedknown, encoded)
        val (success, decoded) = decode64(encoded)
        assertTrue(success)
        assertEquals(originalKnown, decoded)
    }

    @Test
    fun testEncodeDecode() {
        val original = "Texto a codificar"
        val encoded = encode64(original)
        assertFalse(encoded.isEmpty())

        val (success, decoded) = decode64(encoded)
        assertTrue(success)
        assertEquals(original, decoded)
    }


    @Test
    fun testEncodeDecodeEmpty() {
        val original = ""
        val encoded = encode64(original) // should return empty string
        assertTrue(encoded.isEmpty())


        val (success, decoded) = decode64(encoded)
        assertTrue(success)
        assertEquals(original, decoded)
    }



    @Test
    fun testEncodeDecodeInvalid() {
        val original = "Texto a codificar"
        val encoded = encode64(original)
        assertFalse(encoded.isEmpty())

        val (success, decoded) = decode64(encoded + "a")
        assertFalse(success)  // should return false ?
        assertEquals("", decoded)
    }

    @Test
    fun testEncodeDecodeInvalid2() {
        val original = "Texto a codificar"
        val encoded = encode64(original)
        assertFalse(encoded.isEmpty())

        val corrupted = encoded.substring(0, encoded.length - 1) + "a"
        val (success, decoded) = decode64(corrupted)
        assertTrue(success) // should return false ?

        assertNotEquals(original, decoded)
    }

    @Test
    fun textDecodeInvalid() {
        val original = "Texto a codificar"
        val encoded = encode64(original)

        val corrupted = "NOTVALID"
        val (success, decoded) = decode64(corrupted)
        assertTrue(success) // should return false ?

        assertEquals("", decoded) // should return empty string ?
    }

    @Test
    fun testEncodeBytes() {
        val original = "Texto a codificar"
        val tobytes = original.toByteArray()
        val encoded = encode64Bytes(tobytes)
        assertFalse(encoded.isEmpty())

        val (success, decoded) = decode64Bytes(encoded)
        assertTrue(success)

        assertArrayEquals(tobytes, decoded)

    }
}


