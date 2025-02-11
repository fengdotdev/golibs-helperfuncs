#!/usr/bin/env bun

import { randomBytes, createCipheriv, createDecipheriv } from "crypto";

// Typically you would store or derive your key securely.
// Here, we randomly generate a 256-bit key for the demo.
const somekey: Buffer = randomBytes(32);  // AES-256 key (32 bytes)

// For AES-GCM, a 12-byte IV (nonce) is typical.
// NEVER reuse the same key+IV pair for more than one message.
const someiv: Buffer = randomBytes(12);

interface EncryptionResult {
  encryptedBuffer: Buffer;
  authTag: Buffer;
}

class EncodeAESGCMObject {
  data64: string;
  iv64: string;
  authtag64: string;

  constructor(data: Buffer, iv: Buffer, authtag: Buffer) {
    this.data64 = EncodeAESGCMObject.toBase64(data);
    this.iv64 = EncodeAESGCMObject.toBase64(iv);
    this.authtag64 = EncodeAESGCMObject.toBase64(authtag);
  }

  static toBase64(buffer: Buffer): string {
    return buffer.toString("base64");
  }

  static fromBase64(base64String: string): Buffer {
    return Buffer.from(base64String, "base64");
  }

  getData(): Buffer {
    return EncodeAESGCMObject.fromBase64(this.data64);
  }

  getIV(): Buffer {
    return EncodeAESGCMObject.fromBase64(this.iv64);
  }

  getAuthTag(): Buffer {
    return EncodeAESGCMObject.fromBase64(this.authtag64);
  }
}






function encryptAESGCM64(key:Buffer,plaintext: string,iv:Buffer):{ obj:EncodeAESGCMObject,key64:string} {
  // Create cipher
  const cipher = createCipheriv("aes-256-gcm", key, iv);

  // Encrypt the data
  const encryptedBuffer = Buffer.concat([
    cipher.update(plaintext, "utf8"),
    cipher.final()
  ]);

  // GCM produces an authentication tag that we must keep
  const authTag = cipher.getAuthTag();

  return { obj: new EncodeAESGCMObject(encryptedBuffer, iv, authTag), key64: key.toString("base64") };
}


function decryptAESGCM64(Key64:string,obj:EncodeAESGCMObject): string {

 const key = Buffer.from(Key64, "base64");

  // Create decipher
  const decipher = createDecipheriv("aes-256-gcm", key, obj.getIV());

  // Provide the tag that was generated during encryption
  decipher.setAuthTag(obj.getAuthTag());

  // Decrypt the data
  const decryptedBuffer = Buffer.concat([
    decipher.update(obj.getData()),
    decipher.final()
  ]);

  return decryptedBuffer.toString("utf8");
}



async function sendDataToAPI<T>(data: T, key: string, endpoint: string): Promise<string> {
  try {
      const url = new URL(endpoint);
      url.searchParams.append("key", key); // Agrega el query param 'key'

      const response = await fetch(url.toString(), {
          method: "POST",
          headers: {
              "Content-Type": "application/json"
          },
          body: JSON.stringify(data) // Convierte la clase en JSON
      });

      if (!response.ok) {
          throw new Error(`Error: ${response.status} ${response.statusText}`);
      }

      return await response.text(); // Devuelve la respuesta del servidor como texto
  } catch (error) {
      console.error("Error en la solicitud:", error);
      throw error; // Relanza el error para manejarlo externamente si es necesario
  }
}




// Example usage:

console.log("AES-256-GCM Encryption Example");
console.log("Key (hex):        ", somekey.toString("hex"));
console.log("IV (hex):         ", someiv.toString("hex"));
const message = "Hello from Bun with AES-256-GCM!";
console.log("Message:          ", message);

const {obj,key64} = encryptAESGCM64(somekey,message,someiv);

console.log("Encrypted (base64): ", obj.data64);
console.log("Auth Tag (base64):  ", obj.authtag64);
console.log("IV (base64):  ", obj.iv64);
const decryptedMessage64 = decryptAESGCM64(key64, obj);

console.log("Decrypted 64 Text:   ", decryptedMessage64);

const json = JSON.stringify(obj);
console.log("key64: ", key64);
console.log("JSON: ", json);

const result = sendDataToAPI(obj, key64, "http://localhost:8080/encoded")

result.then((data) => {
  console.log("Respuesta del servidor:", data);
}).catch((error) => {
  console.error("Error en la solicitud:", error);
});