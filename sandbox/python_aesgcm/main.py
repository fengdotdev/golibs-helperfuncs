import encrypt

print("hello world")

ciphertext_hex, tag_hex = aes_gcm_encrypt(KEY_HEX, IV_HEX, PT_HEX)
print("Ciphertext =", ciphertext_hex)
print("Tag        =", tag_hex)