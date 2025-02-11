from cryptography.hazmat.primitives.ciphers.aead import AESGCM

def aes_gcm_encrypt(key_hex, iv_hex, plaintext_hex, aad_hex=None):
    key = bytes.fromhex(key_hex)
    iv = bytes.fromhex(iv_hex)
    plaintext = bytes.fromhex(plaintext_hex)
    aad = bytes.fromhex(aad_hex) if aad_hex else None

    aesgcm = AESGCM(key)
    ciphertext_with_tag = aesgcm.encrypt(iv, plaintext, aad)
    # The last 16 bytes are the authentication tag
    ciphertext = ciphertext_with_tag[:-16]
    tag = ciphertext_with_tag[-16:]

    return (ciphertext.hex(), tag.hex())

# --- Test our known vector ---
KEY_HEX = "00000000000000000000000000000000"
IV_HEX  = "000000000000000000000002"
PT_HEX  = "00000000000000000000000000000000"

ciphertext_hex, tag_hex = aes_gcm_encrypt(KEY_HEX, IV_HEX, PT_HEX)
print("Ciphertext =", ciphertext_hex)
print("Tag        =", tag_hex)


