package dev.feng.helperfuncs.commontypes




// This is a simple error type that can be used to return errors from functions golang style

interface Err{
    fun error(): String
}


class SimpleErr(private val error: String): Err{
    override fun error(): String {
        return error
    }
}