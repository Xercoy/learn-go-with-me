Randomize the Elements of a Byte Slice in One Line - Code

View the blog post [here](http://learngowith.me/randomize-the-elements-of-a-byte-slice-in-one-line/)

#### Usage

Run the code sample that randomizes a slice via iteration: `$go run iteration.go`

Sample Output
```
5.08µs - created random slice [217 235 99 200 69 241 191 174 64 73 217 126 38 114 81 54 151 212 188 1 187 97 174 165 243 37 191 2 211 249 172 176 68 159 55 98 81 55 55 85 156 22 97 204 123 161 115 121 102 96 23 24 105 131 168 141 166 80 56 238 17 96 38 195 218 25 173 169 205 235 26 229 186 178 199 194 131 159 144 34 25 218 147 9 20 193 41 11 67 43 156 235 40 10 109 90 64 110 83 244]
```

===

Run the code sample that randomizes a slice via the new `Read()` function: `$go run read.go`

Sample Output
```
1.188µs - created random slice [90 53 181 49 162 205 104 37 192 171 174 192 41 224 35 115 223 112 87 129 84 69 248 171 144 211 83 177 233 215 154 208 132 203 44 216 146 197 59 141 179 245 4 128 105 48 146 205 159 216 157 149 45 49 112 255 237 1 158 29 254 204 171 120 226 44 96 68 182 12 208 153 97 71 204 178 246 156 70 69 58 97 233 123 140 95 179 77 226 177 87 142 182 174 184 193 164 123 223 220]
```