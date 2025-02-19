package dev.feng.helperfuncs


import org.junit.Assert.assertEquals
import org.junit.Assert.fail
import org.junit.Test
import java.util.Base64
import kotlinx.serialization.Serializable


const  private val url = "http://localhost:8080/"




class Helperfuncstest{


    @Serializable
    data class Operation(
        val operation: String,
        val n1: Int,
        val n2: Int
    )

    @Serializable
    data class Result(
        val result: Int
    )



    @Test
    fun TOJsonFromJson(){
        val op = Operation("add",1,2)
        val json = HelperFuncs.toJson(op)
        val op2 = HelperFuncs.fromJson<Operation>(json)
        assertEquals(op,op2)
    }


    @Test
    fun GetResourceT(){
    val endpoint = url+"add"
    }
}