# sed -i '' -f convert.sed *.go
s/\.Float(/\.toFloat64(/g
s/\.Set(/\.set(/g
s/\.Get(/\.get(/g
s/\.Call(/\.call(/g
s/\.Valid(/\.valid(/g
s/\.String(/\.toString(/g
s/\.Int(/\.toInt(/g
s/\.Bool(/\.toBool(/g
s/\.New(/\.jsNew(/g
s/\.ToSlice(/\.toSlice(/g
s/\.Invoke(/\.invoke(/g
