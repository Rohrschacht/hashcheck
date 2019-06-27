# hashcheck

hashcheck is a small tool that verifies that a file matches a given hash.
Given a hashsum and a filepath, hashcheck determines the needed hasher based on the length of the hashsum.
It then calculates the hash of the given file and reports whether the hashes match.

hashcheck uses only the great go standard library. As a result it works on linux, windows and mac os.

I wrote this tool because I frequently check hashes of files I download and am annoyed with having to pick out the right ...sum executable every time.
