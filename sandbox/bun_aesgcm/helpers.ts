
import { randomBytes, createCipheriv, createDecipheriv } from "crypto";

import { AdditionalData, Err, MayAdditionalData, MayErr,Result,SimpleErr } from "./models";


// generate a 256-bit key for the demo.
export function GenerateKey ():Buffer {
    return randomBytes(32);
}

// For AES-GCM, a 12-byte IV (nonce) is typical.
export function GenerateIV ():Buffer {
    return randomBytes(12);
}



export function additionalDataToBuffer(data: AdditionalData): Buffer {
    // Convierte el objeto a cadena JSON y luego a un Buffer
    return Buffer.from(JSON.stringify(data));
  }



export function Encode64(data: Buffer): string {
    return data.toString("base64");
  }

export function Decode64(data: string): Buffer {
    return Buffer.from(data, "base64");
}



export async function sendDataToAPI<T>(data: T, key: string, endpoint: string): Promise<string> {
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

        console.log("Success!");
        return await response.text(); 
    } catch (error) {
        console.error("Error:", error);
        throw error; 
    }
  }



  export function encryptAESGCM(
    key: Buffer,
    iv: Buffer,
    plaintext: Buffer,
    additionalData: MayAdditionalData
  ): Result {
    try {
      // Create cipher
      const cipher = createCipheriv("aes-256-gcm", key, iv);
  
      // Set additional authenticated data (AAD)
      if (additionalData !== null){
        cipher.setAAD(additionalData);
      }
      
      // Encrypt the data
      const encryptedBuffer = Buffer.concat([
        cipher.update(plaintext),
        cipher.final(),
      ]);
  
      // Get the authentication tag
      const authTag = cipher.getAuthTag();
  
      // Option 1: Append the authTag to the ciphertext
      const combinedBuffer = Buffer.concat([encryptedBuffer, authTag]);
  
      return { Data: combinedBuffer, MayErr: null };
    } catch (error) {
      return { Data: Buffer.from(""), MayErr: SimpleErr("failed to encrypt") };
    }
  }
  



  export function decryptAESGCM(
    key: Buffer,
    iv: Buffer,
    cipherTextWithTag: Buffer,
    additionalData: MayAdditionalData
  ): Result {
    try {
      // Assuming the last 16 bytes are the auth tag
      const authTagLength = 16;
      const authTag = cipherTextWithTag.slice(cipherTextWithTag.length - authTagLength);
      const encryptedText = cipherTextWithTag.slice(0, cipherTextWithTag.length - authTagLength);
  
      // Create decipher
      const decipher = createDecipheriv("aes-256-gcm", key, iv);
  

      if (additionalData !== null){
        decipher.setAAD(additionalData);
      }
      // Set AAD
      
  
      // Set the authentication tag that was appended during encryption
      decipher.setAuthTag(authTag);
  
      // Decrypt the data
      const decryptedBuffer = Buffer.concat([
        decipher.update(encryptedText),
        decipher.final(),
      ]);
  
      return { Data: decryptedBuffer, MayErr: null };
    } catch (error) {
      return { Data: Buffer.from(""), MayErr: SimpleErr("failed to decrypt") };
    }
  }