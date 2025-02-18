package dev.feng.helperfuncs.data
import org.junit.Assert.*
import org.junit.Test

class sha256test {


    @Test
    fun TestSha256WithKnownValues() {
        val originalKnown = "foo"
        val encodedknown = "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"

        val encoded = getSHA256(originalKnown)
        assertEquals(encodedknown, encoded)


        val origianlKnown2 = "bar"
        val encodedknown2 = "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9"

        val encoded2 = getSHA256(origianlKnown2)
        assertEquals(encodedknown2, encoded2)
    }

    @Test
    fun TestSha256WithEmpty() {
        val original = ""
        val encoded = getSHA256(original)
        assertEquals("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", encoded)
    }

    @Test
    fun TestSha256BytesWithKnownValues() {
        val originalKnown = "foo"
        val encodedknown = "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"

        val encoded = getSHA256Bytes(originalKnown.toByteArray())
        assertEquals(encodedknown, encoded)
    }
}