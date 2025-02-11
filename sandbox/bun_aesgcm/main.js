#!/usr/bin/env bun

import { randomBytes, createCipheriv, createDecipheriv } from "crypto";

// Typically you would store or derive your key securely.
// Here, we randomly generate a 256-bit key for the demo.
const key = randomBytes(32);  // AES-256 key (32 bytes)

// For AES-GCM, a 12-byte IV (nonce) is typical.
// NEVER reuse the same key+IV pair for more than one message.
const iv = randomBytes(12);

function encryptAESGCM(plaintext) {
  // Create cipher
  const cipher = createCipheriv("aes-256-gcm", key, iv);

  // Encrypt the data
  const encryptedBuffer = Buffer.concat([
    cipher.update(plaintext, "utf8"),
    cipher.final()
  ]);
  
  // GCM produces an authentication tag that we must keep
  const authTag = cipher.getAuthTag();

  return { encryptedBuffer, authTag };
}

function decryptAESGCM(encryptedBuffer, authTag) {
  // Create decipher
  const decipher = createDecipheriv("aes-256-gcm", key, iv);

  // Provide the tag that was generated during encryption
  decipher.setAuthTag(authTag);

  // Decrypt the data
  const decryptedBuffer = Buffer.concat([
    decipher.update(encryptedBuffer),
    decipher.final()
  ]);

  return decryptedBuffer.toString("utf8");
}

// Example usage:
const message = "Hello from Bun with AES-256-GCM!";
const { encryptedBuffer, authTag } = encryptAESGCM(message);

console.log("Encrypted (hex):", encryptedBuffer.toString("hex"));
console.log("Auth Tag (hex):  ", authTag.toString("hex"));

const decryptedMessage = decryptAESGCM(encryptedBuffer, authTag);
console.log("Decrypted Text:   ", decryptedMessage);
