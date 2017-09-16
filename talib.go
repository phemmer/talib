/*
Package talib is a wrapper around the Techinal Analysis Library often used for stock/financial analysis: http://ta-lib.org/

Most of the function arguments are the same as those used by the ta-lib C library. However a few notable points:

outReal - If not provided, a slice of the same size as the input will be generated. If provided, it may point to the same slice as used for input.

Return slice - This will be the same as outReal, but subsliced to remove unused elements.

Return int - This will be the position in the input slice that corresponds to the first element of the output slice.

*/
package talib
