package dev.feng.helperfuncs


import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Assert.fail
import org.junit.Test



const  private val url = "http://localhost:8080/"

class Helperfuncstest{

    data class Operation(
        val operation: String,
        val n1: Int,
        val n2: Int
    )

    data class Result(
        val result: Int
    )
    data class Foo(
        val bar: String
    )



    @Test
    fun TestTOJsonFromJson(){
        val op = Operation("add",1,2)
        val json = HelperFuncs.toJson<Operation> (op)
        val op2 = HelperFuncs.fromJson<Operation>(json)
        assertEquals(op,op2)

        val res = Result(3)
        val json2 = HelperFuncs.toJson<Result>(res)
        val res2 = HelperFuncs.fromJson<Result>(json2)
        assertEquals(res,res2)
    }

    @Test
    fun Test_getResourceAsText(){
        val (text:String,err) = HelperFuncs.getResourceAsText(url+"ping")
        if (err != null) {
            println( "make sure the server is running at $url")
            fail(err.error())
        }
        assertEquals("pong",text)
    }


    @Test
    fun Test_GetResourceJson(){
        val endpoint = url+"add"
        val op = Operation("add",1,2)
        val json = HelperFuncs.toJson<Operation>(op)
        val (data:Result?,err) = HelperFuncs.getResourceWithJson<Result>(endpoint,json)
        if (err != null) {
            println( "make sure the server is running at $url")
            fail(err.error())
        }
        val expectedres = Result(3)
        assertEquals(expectedres,data)
    }

    @Test
    fun Test_GetResourceJsonFail(){
        val endpoint = url+"404"
        val op = Operation("add",1,2)
        val json = HelperFuncs.toJson<Operation>(op)
        val (data:Result?,err) = HelperFuncs.getResourceWithJson<Result>(endpoint,json)
        assertTrue(err != null)
        println("404 expected"+ err?.error())
        val expectedres:Result? = null
        assertEquals(expectedres,data)
    }

    @Test
    fun Test_GetResourceJsonFail2(){
        val endpoint = url+"add"
        val op = Operation("multiply",1,2)
        val json = HelperFuncs.toJson<Operation>(op)
        val (data:Foo?,err) = HelperFuncs.getResourceWithJson<Foo>(endpoint,json)
        assertTrue(err != null)
        println("invalid operation 400"+ err?.error())
        val expectedres:Result? = null
        assertEquals(expectedres,data)
    }

    @Test
    fun GetResourceT(){
    val endpoint = url+"foo"

        val (data:Foo?,err) = HelperFuncs.getResource<Foo>(endpoint)
        if (err != null) {
            println( "make sure the server is running at $url")
            fail(err.error())
        }

        val expectedfoo = Foo("baz")
        assertEquals(expectedfoo,data)
    }




    @Test
    fun GetResourceTFail(){
        val endpoint = url+"404"

        val (data:Foo?,err) = HelperFuncs.getResource<Foo>(endpoint)
        assertTrue(err != null)
        println("404 expected"+ err?.error())
        val expectedfoo:Foo? = null
        assertEquals(expectedfoo,data)
    }
}