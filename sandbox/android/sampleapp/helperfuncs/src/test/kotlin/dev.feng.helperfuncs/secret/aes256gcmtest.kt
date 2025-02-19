package dev.feng.helperfuncs.secret
import org.junit.Assert.assertEquals
import org.junit.Assert.fail
import org.junit.Test
import java.util.Base64

class AES256gcmTest1 {
    val text = "to go server"
    val key64 = "T26jQLcmz49b/UU0exWOblxdEaBlSED96TPlnl89U9k="
    val iv64 = "p2C0G98HwCajYXur"
    val cyphertext64 = "AjIfZOH5FGwvGea9LMCRpal6DjyDRPksZpXJ0A=="
    val additionalData64 = "eyJhbGdvcml0aG0iOiJBRVMiLCJtb2RlIjoiR0NNIiwic3RyZW5ndGgiOjI1NiwiaXY2NCI6InAyQzBHOThId0NhallYdXIifQ=="


    @Test
    fun TestDecryptWithGenerate() {
        val key = generateKey()
        val iv = generateIV()
        val txt = "I met a traveller from an antique land,\n" +
                "Who said—“Two vast and trunkless legs of stone\n" +
                "Stand in the desert. . . . Near them, on the sand,\n" +
                "Half sunk a shattered visage lies, whose frown,\n" +
                "And wrinkled lip, and sneer of cold command,\n" +
                "Tell that its sculptor well those passions read\n" +
                "Which yet survive, stamped on these lifeless things,\n" +
                "The hand that mocked them, and the heart that fed;\n" +
                "And on the pedestal, these words appear:\n" +
                "My name is Ozymandias, King of Kings;\n" +
                "Look on my Works, ye Mighty, and despair!\n" +
                "Nothing beside remains. Round the decay\n" +
                "Of that colossal Wreck, boundless and bare\n" +
                "The lone and level sands stretch far away."
        val additionalData = "some random poem".toByteArray()

        val (cypher:ByteArray,err ) = encodeAESGCM(key, iv, txt.toByteArray(), additionalData)
        if (err != null) {
            fail(err.error())
        }
        val (decoded:ByteArray,err2) = decodeAESGCM(key, iv, cypher, additionalData)
        if (err2 != null) {
            fail(err2.error())
        }
        assertEquals(txt, String(decoded))

    }

    @Test
    fun TestDecryptWithGenerateNoADD() {
        val key = generateKey()
        val iv = generateIV()

        val (cypher:ByteArray,err ) = encodeAESGCM(key, iv, text.toByteArray(), null)
        if (err != null) {
            fail(err.error())
        }
        val (decoded:ByteArray,err2) = decodeAESGCM(key, iv, cypher, null)
        if (err2 != null) {
            fail(err2.error())
        }
        assertEquals(text, String(decoded))

    }
    @Test
    fun TestDecryptWithGenerateNoADD2() {
        val key = generateKey()
        val iv = generateIV()
        val txt = "lorem ipsum dolor sit amet consectetur adipiscing elit"

        val (cypher:ByteArray,err ) = encodeAESGCM(key, iv, txt.toByteArray(), null)
        if (err != null) {
            fail(err.error())
        }
        val (decoded:ByteArray,err2) = decodeAESGCM(key, iv, cypher, null)
        if (err2 != null) {
            fail(err2.error())
        }
        assertEquals(txt, String(decoded))

    }



    @Test
    fun TestEncryptDecryptWithKnows() {

            val key = Base64.getDecoder().decode(key64)
            val iv = Base64.getDecoder().decode(iv64)
            val additionalData = Base64.getDecoder().decode(additionalData64)
            val (cypher:ByteArray,err ) = encodeAESGCM(key, iv, text.toByteArray(), additionalData)
            if (err != null) {
                fail(err.error())
            }
            val cypher64 = Base64.getEncoder().encodeToString(cypher)
            assertEquals(cyphertext64, cypher64)

            val (decoded:ByteArray,err2) = decodeAESGCM(key, iv, cypher, additionalData)
            if (err2 != null) {
                fail(err2.error())
            }

            assertEquals(text, String(decoded))

    }

}