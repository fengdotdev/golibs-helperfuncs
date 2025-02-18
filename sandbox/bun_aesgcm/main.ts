#!/usr/bin/env bun

import { GenerateKey,GenerateIV, encryptAESGCM, Encode64, additionalDataToBuffer, sendDataToAPI, Decode64, decryptAESGCM } from "./helpers";
import { AdditionalData, Payload, Result } from "./models";



async function ToApi(secretmessage: string) {

  console.log("<----------To API---------------->");

  const somekey:Buffer = GenerateKey();
  const somekey64 = Encode64(somekey);

  const someiv: Buffer = GenerateIV();
  const someiv64 = Encode64(someiv);
  
  const plaintext = secretmessage;
  console.log("original data: "+plaintext);
  let data = Buffer.from(plaintext);
  
  const url = "http://localhost:8080/encoded";

  const additionalData: AdditionalData = {
    algorithm: "AES",
    mode: "GCM",
    strength: 256,
    iv64: Encode64(someiv)
  };

  console.log("key64: " + somekey64);
  console.log("iv64: " + someiv64);
  console.log ("data64: " + Encode64(data));
  console.log("additional data: "+ JSON.stringify(additionalData));
  console.log("additional data64: "+ Encode64(additionalDataToBuffer(additionalData)));


  const addbuffer = additionalDataToBuffer(additionalData);

  const result:Result = encryptAESGCM(somekey,someiv,data,addbuffer);

  const payload: Payload = {
    cypher64: Encode64(result.Data),
    additionaldata: additionalData
  };


  const response =await  sendDataToAPI(payload,somekey64,url);


  console.log("Response: ", response);
}



function OwnValues() {

  console.log("<----------Own values---------------->");

  const key = GenerateKey();
  const iv = GenerateIV();
  const plaintext = "Hello, world!";
  console.log("original data: "+plaintext);
  const data = Buffer.from(plaintext);


  //additionalData null

  const result:Result = encryptAESGCM(key,iv,data,null);

  if (result.MayErr) {
    console.error("Error: ", result.MayErr.Error());
  }


  const decryptResult:Result = decryptAESGCM(key,iv,result.Data,null);

  if (decryptResult.MayErr) {
    console.error("Error: ", decryptResult.MayErr.Error());
  }

  console.log("Decrypt Data: ", decryptResult.Data.toString());

  if (decryptResult.Data.toString() === plaintext) {

    console.log("Success!");
  }
  else {
    console.error("Error: Decrypted data does not match plaintext");
  }
}


function KnownValues(){
  console.log("<----------Known values---------------->");
	const text1          = "Hola Mundo"
  const key64_1        = "ZOkodKmzHIMwBI3RtvRlSo4dKQWU5bM3+lKKIvmSy3w="
	const iv64_1         = "32bVr0KW+Cj5pPLB"
	const ciphertext64_1 = "WKBqzxm+x6R2sg5+0e2XLXGpC9QuY68wfiQ="
  console.log("original data: "+text1);


  const key = Decode64(key64_1);
  const iv = Decode64(iv64_1);
  const ciphertext = Decode64(ciphertext64_1);


  const result:Result = decryptAESGCM(key,iv,ciphertext,null);

  if (result.MayErr) {
    console.error("Error: ", result.MayErr.Error());
  }

  console.log("Data: ", result.Data.toString());

  if (result.Data.toString() === text1) {
    console.log("Success!");
  }
  else {
    console.error("Error: Decrypted data does not match plaintext");
  }

}




async function main()  {
  
  OwnValues();
  KnownValues();
  await ToApi("to go server");
}




main();







