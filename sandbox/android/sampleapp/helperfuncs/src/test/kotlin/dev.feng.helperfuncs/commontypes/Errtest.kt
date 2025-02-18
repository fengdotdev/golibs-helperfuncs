package dev.feng.helperfuncs.commontypes

import org.junit.Assert.*
import org.junit.Test



class Errtest {

    @Test
    fun TestSimpleErr() {
        val err = SimpleErr("error")
        assertEquals("error", err.error())
    }

    @Test
    fun TestSimpleErr2() {
        val err = SimpleErr("error2")
        assertEquals("error2", err.error())
    }

    private fun nilerr(): Err? {
        return null
    }
    private fun someerr(): Err {
        return SimpleErr("error")
    }

    @Test
    fun TestSimpleErr3() {

        val result = nilerr()
        assertNull(result)

        val result2 = someerr()
        assertNotNull(result2)
        assertEquals("error", result2.error())
    }
}