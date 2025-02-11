from cryptography.hazmat.primitives.ciphers.aead import AESGCM

def aes_gcm_decrypt(key_hex, iv_hex, ciphertext_hex, tag_hex, aad_hex=None):
    key:bytes = bytes.fromhex(key_hex)
    iv: bytes = bytes.fromhex(iv_hex)
    ciphertext = bytes.fromhex(ciphertext_hex)
    tag = bytes.fromhex(tag_hex)
    aad = bytes.fromhex(aad_hex) if aad_hex else None

    aesgcm = AESGCM(key)
    # GCM expects ciphertext + tag concatenated at decrypt
    plaintext = aesgcm.decrypt(iv, ciphertext + tag, aad)
    return plaintext.hex()

# --- Test with known values ---
KEY_HEX = "00000000000000000000000000000000"
IV_HEX  = "000000000000000000000002"
CT_HEX  = "0388dace60b6a392f328c2b971b2fe78"
TAG_HEX = "ab6e47d42cec13bdf53a67b21257bddf"

plaintext_hex = aes_gcm_decrypt(KEY_HEX, IV_HEX, CT_HEX, TAG_HEX)
print("Plaintext =", plaintext_hex)
