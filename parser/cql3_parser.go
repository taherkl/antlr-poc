// Generated from CQL3.g4 by ANTLR 4.7.

package parser // CQL3

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 92, 737,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 3, 2, 3, 2, 6, 2, 137, 10, 2, 13, 2, 14, 2,
	138, 3, 2, 7, 2, 142, 10, 2, 12, 2, 14, 2, 145, 11, 2, 3, 2, 6, 2, 148,
	10, 2, 13, 2, 14, 2, 149, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 166, 10, 3, 3, 4, 3, 4, 6,
	4, 170, 10, 4, 13, 4, 14, 4, 171, 3, 4, 7, 4, 175, 10, 4, 12, 4, 14, 4,
	178, 11, 4, 3, 4, 6, 4, 181, 10, 4, 13, 4, 14, 4, 182, 3, 5, 3, 5, 3, 5,
	5, 5, 188, 10, 5, 3, 6, 3, 6, 3, 6, 5, 6, 193, 10, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 208,
	10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 218, 10,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 224, 10, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 244, 10, 12, 3, 13, 3, 13, 3, 13,
	5, 13, 249, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5,
	15, 258, 10, 15, 3, 15, 3, 15, 5, 15, 262, 10, 15, 3, 15, 5, 15, 265, 10,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15,
	276, 10, 15, 5, 15, 278, 10, 15, 3, 16, 3, 16, 3, 16, 5, 16, 283, 10, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 294,
	10, 17, 3, 17, 5, 17, 297, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 303,
	10, 18, 12, 18, 14, 18, 306, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 7, 19, 314, 10, 19, 12, 19, 14, 19, 317, 11, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 7, 20, 325, 10, 20, 12, 20, 14, 20, 328, 11, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 334, 10, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 347, 10, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 354, 10, 25, 3, 26, 3, 26, 3,
	26, 7, 26, 359, 10, 26, 12, 26, 14, 26, 362, 11, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 375, 10,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 5, 27, 390, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7,
	28, 396, 10, 28, 12, 28, 14, 28, 399, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 411, 10, 29, 3, 30, 3,
	30, 3, 30, 7, 30, 416, 10, 30, 12, 30, 14, 30, 419, 11, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 431, 10,
	31, 12, 31, 14, 31, 434, 11, 31, 5, 31, 436, 10, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 5, 31, 444, 10, 31, 3, 32, 3, 32, 5, 32, 448, 10,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 455, 10, 32, 3, 32, 3, 32,
	3, 32, 5, 32, 460, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 467,
	10, 33, 12, 33, 14, 33, 470, 11, 33, 5, 33, 472, 10, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 5, 34, 479, 10, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 7, 35, 487, 10, 35, 12, 35, 14, 35, 490, 11, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 5, 36, 497, 10, 36, 3, 37, 3, 37, 5, 37, 501, 10,
	37, 3, 37, 3, 37, 5, 37, 505, 10, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38,
	3, 38, 3, 38, 3, 38, 7, 38, 515, 10, 38, 12, 38, 14, 38, 518, 11, 38, 3,
	39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 5, 40, 526, 10, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 7, 42, 535, 10, 42, 12, 42, 14, 42,
	538, 11, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 545, 10, 43, 3,
	44, 3, 44, 3, 44, 3, 44, 7, 44, 551, 10, 44, 12, 44, 14, 44, 554, 11, 44,
	3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 5, 45, 561, 10, 45, 3, 45, 3, 45, 5,
	45, 565, 10, 45, 3, 45, 3, 45, 3, 45, 5, 45, 570, 10, 45, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 47, 3, 47, 7, 47, 578, 10, 47, 12, 47, 14, 47, 581, 11,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 7, 48, 590, 10, 48,
	12, 48, 14, 48, 593, 11, 48, 3, 48, 3, 48, 5, 48, 597, 10, 48, 3, 49, 3,
	49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 5, 52, 614, 10, 52, 3, 53, 3, 53, 3, 53, 5, 53, 619,
	10, 53, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 625, 10, 54, 3, 55, 3, 55, 3,
	55, 5, 55, 630, 10, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 7, 56, 641, 10, 56, 12, 56, 14, 56, 644, 11, 56, 5, 56, 646,
	10, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 7, 57, 654, 10, 57, 12,
	57, 14, 57, 657, 11, 57, 5, 57, 659, 10, 57, 3, 57, 3, 57, 3, 58, 3, 58,
	3, 58, 3, 58, 7, 58, 667, 10, 58, 12, 58, 14, 58, 670, 11, 58, 5, 58, 672,
	10, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 7, 59, 681, 10,
	59, 12, 59, 14, 59, 684, 11, 59, 5, 59, 686, 10, 59, 3, 59, 3, 59, 3, 60,
	3, 60, 3, 60, 7, 60, 693, 10, 60, 12, 60, 14, 60, 696, 11, 60, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 5, 63, 707, 10, 63,
	3, 64, 3, 64, 3, 64, 5, 64, 712, 10, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3,
	66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 733, 10, 66, 3, 67, 3, 67, 3, 67, 2,
	2, 68, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 2,
	7, 4, 2, 45, 45, 71, 71, 3, 2, 8, 9, 4, 2, 47, 47, 77, 77, 3, 2, 17, 32,
	4, 2, 53, 53, 73, 73, 2, 774, 2, 134, 3, 2, 2, 2, 4, 165, 3, 2, 2, 2, 6,
	167, 3, 2, 2, 2, 8, 187, 3, 2, 2, 2, 10, 189, 3, 2, 2, 2, 12, 198, 3, 2,
	2, 2, 14, 204, 3, 2, 2, 2, 16, 211, 3, 2, 2, 2, 18, 214, 3, 2, 2, 2, 20,
	225, 3, 2, 2, 2, 22, 243, 3, 2, 2, 2, 24, 245, 3, 2, 2, 2, 26, 252, 3,
	2, 2, 2, 28, 255, 3, 2, 2, 2, 30, 279, 3, 2, 2, 2, 32, 286, 3, 2, 2, 2,
	34, 298, 3, 2, 2, 2, 36, 309, 3, 2, 2, 2, 38, 320, 3, 2, 2, 2, 40, 333,
	3, 2, 2, 2, 42, 335, 3, 2, 2, 2, 44, 337, 3, 2, 2, 2, 46, 339, 3, 2, 2,
	2, 48, 343, 3, 2, 2, 2, 50, 355, 3, 2, 2, 2, 52, 389, 3, 2, 2, 2, 54, 391,
	3, 2, 2, 2, 56, 410, 3, 2, 2, 2, 58, 412, 3, 2, 2, 2, 60, 443, 3, 2, 2,
	2, 62, 445, 3, 2, 2, 2, 64, 461, 3, 2, 2, 2, 66, 473, 3, 2, 2, 2, 68, 483,
	3, 2, 2, 2, 70, 491, 3, 2, 2, 2, 72, 498, 3, 2, 2, 2, 74, 510, 3, 2, 2,
	2, 76, 519, 3, 2, 2, 2, 78, 525, 3, 2, 2, 2, 80, 529, 3, 2, 2, 2, 82, 531,
	3, 2, 2, 2, 84, 544, 3, 2, 2, 2, 86, 546, 3, 2, 2, 2, 88, 569, 3, 2, 2,
	2, 90, 571, 3, 2, 2, 2, 92, 573, 3, 2, 2, 2, 94, 596, 3, 2, 2, 2, 96, 598,
	3, 2, 2, 2, 98, 600, 3, 2, 2, 2, 100, 604, 3, 2, 2, 2, 102, 613, 3, 2,
	2, 2, 104, 618, 3, 2, 2, 2, 106, 624, 3, 2, 2, 2, 108, 629, 3, 2, 2, 2,
	110, 631, 3, 2, 2, 2, 112, 649, 3, 2, 2, 2, 114, 662, 3, 2, 2, 2, 116,
	675, 3, 2, 2, 2, 118, 689, 3, 2, 2, 2, 120, 697, 3, 2, 2, 2, 122, 701,
	3, 2, 2, 2, 124, 706, 3, 2, 2, 2, 126, 711, 3, 2, 2, 2, 128, 713, 3, 2,
	2, 2, 130, 732, 3, 2, 2, 2, 132, 734, 3, 2, 2, 2, 134, 143, 5, 4, 3, 2,
	135, 137, 7, 3, 2, 2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142,
	5, 4, 3, 2, 141, 136, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2,
	2, 2, 143, 144, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2,
	146, 148, 7, 3, 2, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 3, 3, 2, 2, 2, 151, 166, 5,
	14, 8, 2, 152, 166, 5, 10, 6, 2, 153, 166, 5, 12, 7, 2, 154, 166, 5, 16,
	9, 2, 155, 166, 5, 18, 10, 2, 156, 166, 5, 20, 11, 2, 157, 166, 5, 24,
	13, 2, 158, 166, 5, 26, 14, 2, 159, 166, 5, 28, 15, 2, 160, 166, 5, 30,
	16, 2, 161, 166, 5, 32, 17, 2, 162, 166, 5, 48, 25, 2, 163, 166, 5, 62,
	32, 2, 164, 166, 5, 72, 37, 2, 165, 151, 3, 2, 2, 2, 165, 152, 3, 2, 2,
	2, 165, 153, 3, 2, 2, 2, 165, 154, 3, 2, 2, 2, 165, 155, 3, 2, 2, 2, 165,
	156, 3, 2, 2, 2, 165, 157, 3, 2, 2, 2, 165, 158, 3, 2, 2, 2, 165, 159,
	3, 2, 2, 2, 165, 160, 3, 2, 2, 2, 165, 161, 3, 2, 2, 2, 165, 162, 3, 2,
	2, 2, 165, 163, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2, 166, 5, 3, 2, 2, 2, 167,
	176, 5, 8, 5, 2, 168, 170, 7, 3, 2, 2, 169, 168, 3, 2, 2, 2, 170, 171,
	3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 173, 3, 2,
	2, 2, 173, 175, 5, 8, 5, 2, 174, 169, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2,
	176, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178,
	176, 3, 2, 2, 2, 179, 181, 7, 3, 2, 2, 180, 179, 3, 2, 2, 2, 181, 182,
	3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 7, 3, 2, 2,
	2, 184, 188, 5, 32, 17, 2, 185, 188, 5, 48, 25, 2, 186, 188, 5, 62, 32,
	2, 187, 184, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 186, 3, 2, 2, 2, 188,
	9, 3, 2, 2, 2, 189, 190, 7, 48, 2, 2, 190, 192, 7, 61, 2, 2, 191, 193,
	5, 98, 50, 2, 192, 191, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3,
	2, 2, 2, 194, 195, 5, 96, 49, 2, 195, 196, 7, 83, 2, 2, 196, 197, 5, 118,
	60, 2, 197, 11, 3, 2, 2, 2, 198, 199, 7, 39, 2, 2, 199, 200, 7, 61, 2,
	2, 200, 201, 5, 96, 49, 2, 201, 202, 7, 83, 2, 2, 202, 203, 5, 118, 60,
	2, 203, 13, 3, 2, 2, 2, 204, 205, 7, 51, 2, 2, 205, 207, 7, 61, 2, 2, 206,
	208, 5, 100, 51, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209,
	3, 2, 2, 2, 209, 210, 5, 96, 49, 2, 210, 15, 3, 2, 2, 2, 211, 212, 7, 79,
	2, 2, 212, 213, 5, 96, 49, 2, 213, 17, 3, 2, 2, 2, 214, 215, 7, 48, 2,
	2, 215, 217, 9, 2, 2, 2, 216, 218, 5, 98, 50, 2, 217, 216, 3, 2, 2, 2,
	217, 218, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 5, 78, 40, 2, 220,
	223, 5, 86, 44, 2, 221, 222, 7, 83, 2, 2, 222, 224, 5, 82, 42, 2, 223,
	221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 19, 3, 2, 2, 2, 225, 226, 7,
	39, 2, 2, 226, 227, 9, 2, 2, 2, 227, 228, 5, 78, 40, 2, 228, 229, 5, 22,
	12, 2, 229, 21, 3, 2, 2, 2, 230, 231, 7, 39, 2, 2, 231, 232, 5, 80, 41,
	2, 232, 233, 7, 76, 2, 2, 233, 234, 5, 90, 46, 2, 234, 244, 3, 2, 2, 2,
	235, 236, 7, 38, 2, 2, 236, 237, 5, 80, 41, 2, 237, 238, 5, 90, 46, 2,
	238, 244, 3, 2, 2, 2, 239, 240, 7, 51, 2, 2, 240, 244, 5, 80, 41, 2, 241,
	242, 7, 83, 2, 2, 242, 244, 5, 82, 42, 2, 243, 230, 3, 2, 2, 2, 243, 235,
	3, 2, 2, 2, 243, 239, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 244, 23, 3, 2,
	2, 2, 245, 246, 7, 51, 2, 2, 246, 248, 7, 71, 2, 2, 247, 249, 5, 100, 51,
	2, 248, 247, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250,
	251, 5, 78, 40, 2, 251, 25, 3, 2, 2, 2, 252, 253, 7, 74, 2, 2, 253, 254,
	5, 78, 40, 2, 254, 27, 3, 2, 2, 2, 255, 257, 7, 48, 2, 2, 256, 258, 7,
	49, 2, 2, 257, 256, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259, 3, 2, 2,
	2, 259, 261, 7, 57, 2, 2, 260, 262, 5, 98, 50, 2, 261, 260, 3, 2, 2, 2,
	261, 262, 3, 2, 2, 2, 262, 264, 3, 2, 2, 2, 263, 265, 5, 42, 22, 2, 264,
	263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267,
	7, 63, 2, 2, 267, 268, 5, 78, 40, 2, 268, 269, 7, 4, 2, 2, 269, 270, 5,
	80, 41, 2, 270, 277, 7, 5, 2, 2, 271, 272, 7, 80, 2, 2, 272, 275, 5, 44,
	23, 2, 273, 274, 7, 83, 2, 2, 274, 276, 5, 46, 24, 2, 275, 273, 3, 2, 2,
	2, 275, 276, 3, 2, 2, 2, 276, 278, 3, 2, 2, 2, 277, 271, 3, 2, 2, 2, 277,
	278, 3, 2, 2, 2, 278, 29, 3, 2, 2, 2, 279, 280, 7, 51, 2, 2, 280, 282,
	7, 57, 2, 2, 281, 283, 5, 100, 51, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3,
	2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 285, 5, 42, 22, 2, 285, 31, 3, 2, 2,
	2, 286, 287, 7, 58, 2, 2, 287, 288, 7, 59, 2, 2, 288, 289, 5, 78, 40, 2,
	289, 290, 5, 34, 18, 2, 290, 291, 7, 81, 2, 2, 291, 293, 5, 36, 19, 2,
	292, 294, 5, 98, 50, 2, 293, 292, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294,
	296, 3, 2, 2, 2, 295, 297, 5, 38, 20, 2, 296, 295, 3, 2, 2, 2, 296, 297,
	3, 2, 2, 2, 297, 33, 3, 2, 2, 2, 298, 299, 7, 4, 2, 2, 299, 304, 5, 80,
	41, 2, 300, 301, 7, 6, 2, 2, 301, 303, 5, 80, 41, 2, 302, 300, 3, 2, 2,
	2, 303, 306, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305,
	307, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 307, 308, 7, 5, 2, 2, 308, 35, 3,
	2, 2, 2, 309, 310, 7, 4, 2, 2, 310, 315, 5, 106, 54, 2, 311, 312, 7, 6,
	2, 2, 312, 314, 5, 106, 54, 2, 313, 311, 3, 2, 2, 2, 314, 317, 3, 2, 2,
	2, 315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3, 2, 2, 2, 317,
	315, 3, 2, 2, 2, 318, 319, 7, 5, 2, 2, 319, 37, 3, 2, 2, 2, 320, 321, 7,
	80, 2, 2, 321, 326, 5, 40, 21, 2, 322, 323, 7, 40, 2, 2, 323, 325, 5, 40,
	21, 2, 324, 322, 3, 2, 2, 2, 325, 328, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2,
	326, 327, 3, 2, 2, 2, 327, 39, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 329, 330,
	7, 72, 2, 2, 330, 334, 7, 86, 2, 2, 331, 332, 7, 75, 2, 2, 332, 334, 7,
	86, 2, 2, 333, 329, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2, 334, 41, 3, 2, 2,
	2, 335, 336, 7, 84, 2, 2, 336, 43, 3, 2, 2, 2, 337, 338, 7, 85, 2, 2, 338,
	45, 3, 2, 2, 2, 339, 340, 7, 64, 2, 2, 340, 341, 7, 7, 2, 2, 341, 342,
	5, 110, 56, 2, 342, 47, 3, 2, 2, 2, 343, 344, 7, 78, 2, 2, 344, 346, 5,
	78, 40, 2, 345, 347, 5, 38, 20, 2, 346, 345, 3, 2, 2, 2, 346, 347, 3, 2,
	2, 2, 347, 348, 3, 2, 2, 2, 348, 349, 7, 68, 2, 2, 349, 350, 5, 50, 26,
	2, 350, 351, 7, 82, 2, 2, 351, 353, 5, 58, 30, 2, 352, 354, 5, 54, 28,
	2, 353, 352, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 49, 3, 2, 2, 2, 355,
	360, 5, 52, 27, 2, 356, 357, 7, 6, 2, 2, 357, 359, 5, 52, 27, 2, 358, 356,
	3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2,
	2, 2, 361, 51, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 364, 5, 80, 41, 2,
	364, 365, 7, 7, 2, 2, 365, 366, 5, 106, 54, 2, 366, 390, 3, 2, 2, 2, 367,
	368, 5, 80, 41, 2, 368, 369, 7, 7, 2, 2, 369, 370, 5, 80, 41, 2, 370, 374,
	9, 3, 2, 2, 371, 375, 7, 86, 2, 2, 372, 375, 5, 112, 57, 2, 373, 375, 5,
	114, 58, 2, 374, 371, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 373, 3, 2,
	2, 2, 375, 390, 3, 2, 2, 2, 376, 377, 5, 80, 41, 2, 377, 378, 7, 7, 2,
	2, 378, 379, 5, 80, 41, 2, 379, 380, 7, 8, 2, 2, 380, 381, 5, 110, 56,
	2, 381, 390, 3, 2, 2, 2, 382, 383, 5, 80, 41, 2, 383, 384, 7, 10, 2, 2,
	384, 385, 5, 106, 54, 2, 385, 386, 7, 11, 2, 2, 386, 387, 7, 7, 2, 2, 387,
	388, 5, 106, 54, 2, 388, 390, 3, 2, 2, 2, 389, 363, 3, 2, 2, 2, 389, 367,
	3, 2, 2, 2, 389, 376, 3, 2, 2, 2, 389, 382, 3, 2, 2, 2, 390, 53, 3, 2,
	2, 2, 391, 392, 7, 55, 2, 2, 392, 397, 5, 56, 29, 2, 393, 394, 7, 40, 2,
	2, 394, 396, 5, 56, 29, 2, 395, 393, 3, 2, 2, 2, 396, 399, 3, 2, 2, 2,
	397, 395, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 55, 3, 2, 2, 2, 399, 397,
	3, 2, 2, 2, 400, 401, 7, 84, 2, 2, 401, 402, 7, 7, 2, 2, 402, 411, 5, 106,
	54, 2, 403, 404, 7, 84, 2, 2, 404, 405, 7, 10, 2, 2, 405, 406, 5, 106,
	54, 2, 406, 407, 7, 11, 2, 2, 407, 408, 7, 7, 2, 2, 408, 409, 5, 106, 54,
	2, 409, 411, 3, 2, 2, 2, 410, 400, 3, 2, 2, 2, 410, 403, 3, 2, 2, 2, 411,
	57, 3, 2, 2, 2, 412, 417, 5, 60, 31, 2, 413, 414, 7, 40, 2, 2, 414, 416,
	5, 60, 31, 2, 415, 413, 3, 2, 2, 2, 416, 419, 3, 2, 2, 2, 417, 415, 3,
	2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 59, 3, 2, 2, 2, 419, 417, 3, 2, 2,
	2, 420, 421, 5, 80, 41, 2, 421, 422, 7, 7, 2, 2, 422, 423, 5, 106, 54,
	2, 423, 444, 3, 2, 2, 2, 424, 425, 5, 80, 41, 2, 425, 426, 7, 56, 2, 2,
	426, 435, 7, 4, 2, 2, 427, 432, 5, 106, 54, 2, 428, 429, 7, 6, 2, 2, 429,
	431, 5, 106, 54, 2, 430, 428, 3, 2, 2, 2, 431, 434, 3, 2, 2, 2, 432, 430,
	3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 436, 3, 2, 2, 2, 434, 432, 3, 2,
	2, 2, 435, 427, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2,
	437, 438, 7, 5, 2, 2, 438, 444, 3, 2, 2, 2, 439, 440, 5, 80, 41, 2, 440,
	441, 7, 56, 2, 2, 441, 442, 7, 12, 2, 2, 442, 444, 3, 2, 2, 2, 443, 420,
	3, 2, 2, 2, 443, 424, 3, 2, 2, 2, 443, 439, 3, 2, 2, 2, 444, 61, 3, 2,
	2, 2, 445, 447, 7, 50, 2, 2, 446, 448, 5, 68, 35, 2, 447, 446, 3, 2, 2,
	2, 447, 448, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 450, 7, 54, 2, 2, 450,
	454, 5, 78, 40, 2, 451, 452, 7, 80, 2, 2, 452, 453, 7, 72, 2, 2, 453, 455,
	7, 86, 2, 2, 454, 451, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 3, 2,
	2, 2, 456, 457, 7, 82, 2, 2, 457, 459, 5, 58, 30, 2, 458, 460, 5, 64, 33,
	2, 459, 458, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 63, 3, 2, 2, 2, 461,
	471, 7, 55, 2, 2, 462, 472, 7, 52, 2, 2, 463, 468, 5, 66, 34, 2, 464, 465,
	7, 40, 2, 2, 465, 467, 5, 66, 34, 2, 466, 464, 3, 2, 2, 2, 467, 470, 3,
	2, 2, 2, 468, 466, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 472, 3, 2, 2,
	2, 470, 468, 3, 2, 2, 2, 471, 462, 3, 2, 2, 2, 471, 463, 3, 2, 2, 2, 472,
	65, 3, 2, 2, 2, 473, 478, 7, 84, 2, 2, 474, 475, 7, 10, 2, 2, 475, 476,
	5, 106, 54, 2, 476, 477, 7, 11, 2, 2, 477, 479, 3, 2, 2, 2, 478, 474, 3,
	2, 2, 2, 478, 479, 3, 2, 2, 2, 479, 480, 3, 2, 2, 2, 480, 481, 7, 7, 2,
	2, 481, 482, 5, 106, 54, 2, 482, 67, 3, 2, 2, 2, 483, 488, 5, 70, 36, 2,
	484, 485, 7, 6, 2, 2, 485, 487, 5, 70, 36, 2, 486, 484, 3, 2, 2, 2, 487,
	490, 3, 2, 2, 2, 488, 486, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489, 69, 3,
	2, 2, 2, 490, 488, 3, 2, 2, 2, 491, 496, 7, 84, 2, 2, 492, 493, 7, 10,
	2, 2, 493, 494, 5, 106, 54, 2, 494, 495, 7, 11, 2, 2, 495, 497, 3, 2, 2,
	2, 496, 492, 3, 2, 2, 2, 496, 497, 3, 2, 2, 2, 497, 71, 3, 2, 2, 2, 498,
	500, 7, 43, 2, 2, 499, 501, 9, 4, 2, 2, 500, 499, 3, 2, 2, 2, 500, 501,
	3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 504, 7, 42, 2, 2, 503, 505, 5, 74,
	38, 2, 504, 503, 3, 2, 2, 2, 504, 505, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2,
	506, 507, 5, 6, 4, 2, 507, 508, 7, 41, 2, 2, 508, 509, 7, 42, 2, 2, 509,
	73, 3, 2, 2, 2, 510, 511, 7, 80, 2, 2, 511, 516, 5, 76, 39, 2, 512, 513,
	7, 40, 2, 2, 513, 515, 5, 76, 39, 2, 514, 512, 3, 2, 2, 2, 515, 518, 3,
	2, 2, 2, 516, 514, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 75, 3, 2, 2,
	2, 518, 516, 3, 2, 2, 2, 519, 520, 7, 72, 2, 2, 520, 521, 7, 86, 2, 2,
	521, 77, 3, 2, 2, 2, 522, 523, 5, 96, 49, 2, 523, 524, 7, 13, 2, 2, 524,
	526, 3, 2, 2, 2, 525, 522, 3, 2, 2, 2, 525, 526, 3, 2, 2, 2, 526, 527,
	3, 2, 2, 2, 527, 528, 7, 84, 2, 2, 528, 79, 3, 2, 2, 2, 529, 530, 7, 84,
	2, 2, 530, 81, 3, 2, 2, 2, 531, 536, 5, 84, 43, 2, 532, 533, 7, 40, 2,
	2, 533, 535, 5, 84, 43, 2, 534, 532, 3, 2, 2, 2, 535, 538, 3, 2, 2, 2,
	536, 534, 3, 2, 2, 2, 536, 537, 3, 2, 2, 2, 537, 83, 3, 2, 2, 2, 538, 536,
	3, 2, 2, 2, 539, 545, 5, 120, 61, 2, 540, 541, 7, 46, 2, 2, 541, 545, 7,
	70, 2, 2, 542, 543, 7, 44, 2, 2, 543, 545, 7, 65, 2, 2, 544, 539, 3, 2,
	2, 2, 544, 540, 3, 2, 2, 2, 544, 542, 3, 2, 2, 2, 545, 85, 3, 2, 2, 2,
	546, 547, 7, 4, 2, 2, 547, 552, 5, 88, 45, 2, 548, 549, 7, 6, 2, 2, 549,
	551, 5, 88, 45, 2, 550, 548, 3, 2, 2, 2, 551, 554, 3, 2, 2, 2, 552, 550,
	3, 2, 2, 2, 552, 553, 3, 2, 2, 2, 553, 555, 3, 2, 2, 2, 554, 552, 3, 2,
	2, 2, 555, 556, 7, 5, 2, 2, 556, 87, 3, 2, 2, 2, 557, 558, 5, 80, 41, 2,
	558, 560, 5, 90, 46, 2, 559, 561, 7, 69, 2, 2, 560, 559, 3, 2, 2, 2, 560,
	561, 3, 2, 2, 2, 561, 564, 3, 2, 2, 2, 562, 563, 7, 66, 2, 2, 563, 565,
	7, 60, 2, 2, 564, 562, 3, 2, 2, 2, 564, 565, 3, 2, 2, 2, 565, 570, 3, 2,
	2, 2, 566, 567, 7, 66, 2, 2, 567, 568, 7, 60, 2, 2, 568, 570, 5, 92, 47,
	2, 569, 557, 3, 2, 2, 2, 569, 566, 3, 2, 2, 2, 570, 89, 3, 2, 2, 2, 571,
	572, 5, 126, 64, 2, 572, 91, 3, 2, 2, 2, 573, 574, 7, 4, 2, 2, 574, 579,
	5, 94, 48, 2, 575, 576, 7, 6, 2, 2, 576, 578, 5, 80, 41, 2, 577, 575, 3,
	2, 2, 2, 578, 581, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 579, 580, 3, 2, 2,
	2, 580, 582, 3, 2, 2, 2, 581, 579, 3, 2, 2, 2, 582, 583, 7, 5, 2, 2, 583,
	93, 3, 2, 2, 2, 584, 597, 5, 80, 41, 2, 585, 586, 7, 4, 2, 2, 586, 591,
	5, 80, 41, 2, 587, 588, 7, 6, 2, 2, 588, 590, 5, 80, 41, 2, 589, 587, 3,
	2, 2, 2, 590, 593, 3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 591, 592, 3, 2, 2,
	2, 592, 594, 3, 2, 2, 2, 593, 591, 3, 2, 2, 2, 594, 595, 7, 5, 2, 2, 595,
	597, 3, 2, 2, 2, 596, 584, 3, 2, 2, 2, 596, 585, 3, 2, 2, 2, 597, 95, 3,
	2, 2, 2, 598, 599, 7, 84, 2, 2, 599, 97, 3, 2, 2, 2, 600, 601, 7, 55, 2,
	2, 601, 602, 7, 62, 2, 2, 602, 603, 7, 52, 2, 2, 603, 99, 3, 2, 2, 2, 604,
	605, 7, 55, 2, 2, 605, 606, 7, 52, 2, 2, 606, 101, 3, 2, 2, 2, 607, 614,
	7, 85, 2, 2, 608, 614, 7, 86, 2, 2, 609, 614, 7, 87, 2, 2, 610, 614, 5,
	132, 67, 2, 611, 614, 7, 88, 2, 2, 612, 614, 7, 89, 2, 2, 613, 607, 3,
	2, 2, 2, 613, 608, 3, 2, 2, 2, 613, 609, 3, 2, 2, 2, 613, 610, 3, 2, 2,
	2, 613, 611, 3, 2, 2, 2, 613, 612, 3, 2, 2, 2, 614, 103, 3, 2, 2, 2, 615,
	619, 7, 12, 2, 2, 616, 617, 7, 14, 2, 2, 617, 619, 7, 84, 2, 2, 618, 615,
	3, 2, 2, 2, 618, 616, 3, 2, 2, 2, 619, 105, 3, 2, 2, 2, 620, 625, 5, 102,
	52, 2, 621, 625, 5, 108, 55, 2, 622, 625, 5, 104, 53, 2, 623, 625, 5, 116,
	59, 2, 624, 620, 3, 2, 2, 2, 624, 621, 3, 2, 2, 2, 624, 622, 3, 2, 2, 2,
	624, 623, 3, 2, 2, 2, 625, 107, 3, 2, 2, 2, 626, 630, 5, 110, 56, 2, 627,
	630, 5, 112, 57, 2, 628, 630, 5, 114, 58, 2, 629, 626, 3, 2, 2, 2, 629,
	627, 3, 2, 2, 2, 629, 628, 3, 2, 2, 2, 630, 109, 3, 2, 2, 2, 631, 645,
	7, 15, 2, 2, 632, 633, 5, 106, 54, 2, 633, 634, 7, 14, 2, 2, 634, 642,
	5, 106, 54, 2, 635, 636, 7, 6, 2, 2, 636, 637, 5, 106, 54, 2, 637, 638,
	7, 14, 2, 2, 638, 639, 5, 106, 54, 2, 639, 641, 3, 2, 2, 2, 640, 635, 3,
	2, 2, 2, 641, 644, 3, 2, 2, 2, 642, 640, 3, 2, 2, 2, 642, 643, 3, 2, 2,
	2, 643, 646, 3, 2, 2, 2, 644, 642, 3, 2, 2, 2, 645, 632, 3, 2, 2, 2, 645,
	646, 3, 2, 2, 2, 646, 647, 3, 2, 2, 2, 647, 648, 7, 16, 2, 2, 648, 111,
	3, 2, 2, 2, 649, 658, 7, 15, 2, 2, 650, 655, 5, 106, 54, 2, 651, 652, 7,
	6, 2, 2, 652, 654, 5, 106, 54, 2, 653, 651, 3, 2, 2, 2, 654, 657, 3, 2,
	2, 2, 655, 653, 3, 2, 2, 2, 655, 656, 3, 2, 2, 2, 656, 659, 3, 2, 2, 2,
	657, 655, 3, 2, 2, 2, 658, 650, 3, 2, 2, 2, 658, 659, 3, 2, 2, 2, 659,
	660, 3, 2, 2, 2, 660, 661, 7, 16, 2, 2, 661, 113, 3, 2, 2, 2, 662, 671,
	7, 10, 2, 2, 663, 668, 5, 106, 54, 2, 664, 665, 7, 6, 2, 2, 665, 667, 5,
	106, 54, 2, 666, 664, 3, 2, 2, 2, 667, 670, 3, 2, 2, 2, 668, 666, 3, 2,
	2, 2, 668, 669, 3, 2, 2, 2, 669, 672, 3, 2, 2, 2, 670, 668, 3, 2, 2, 2,
	671, 663, 3, 2, 2, 2, 671, 672, 3, 2, 2, 2, 672, 673, 3, 2, 2, 2, 673,
	674, 7, 11, 2, 2, 674, 115, 3, 2, 2, 2, 675, 676, 7, 84, 2, 2, 676, 685,
	7, 4, 2, 2, 677, 682, 5, 106, 54, 2, 678, 679, 7, 6, 2, 2, 679, 681, 5,
	106, 54, 2, 680, 678, 3, 2, 2, 2, 681, 684, 3, 2, 2, 2, 682, 680, 3, 2,
	2, 2, 682, 683, 3, 2, 2, 2, 683, 686, 3, 2, 2, 2, 684, 682, 3, 2, 2, 2,
	685, 677, 3, 2, 2, 2, 685, 686, 3, 2, 2, 2, 686, 687, 3, 2, 2, 2, 687,
	688, 7, 5, 2, 2, 688, 117, 3, 2, 2, 2, 689, 694, 5, 120, 61, 2, 690, 691,
	7, 40, 2, 2, 691, 693, 5, 120, 61, 2, 692, 690, 3, 2, 2, 2, 693, 696, 3,
	2, 2, 2, 694, 692, 3, 2, 2, 2, 694, 695, 3, 2, 2, 2, 695, 119, 3, 2, 2,
	2, 696, 694, 3, 2, 2, 2, 697, 698, 5, 122, 62, 2, 698, 699, 7, 7, 2, 2,
	699, 700, 5, 124, 63, 2, 700, 121, 3, 2, 2, 2, 701, 702, 7, 84, 2, 2, 702,
	123, 3, 2, 2, 2, 703, 707, 7, 84, 2, 2, 704, 707, 5, 102, 52, 2, 705, 707,
	5, 110, 56, 2, 706, 703, 3, 2, 2, 2, 706, 704, 3, 2, 2, 2, 706, 705, 3,
	2, 2, 2, 707, 125, 3, 2, 2, 2, 708, 712, 5, 128, 65, 2, 709, 712, 5, 130,
	66, 2, 710, 712, 7, 85, 2, 2, 711, 708, 3, 2, 2, 2, 711, 709, 3, 2, 2,
	2, 711, 710, 3, 2, 2, 2, 712, 127, 3, 2, 2, 2, 713, 714, 9, 5, 2, 2, 714,
	129, 3, 2, 2, 2, 715, 716, 7, 33, 2, 2, 716, 717, 7, 34, 2, 2, 717, 718,
	5, 128, 65, 2, 718, 719, 7, 35, 2, 2, 719, 733, 3, 2, 2, 2, 720, 721, 7,
	36, 2, 2, 721, 722, 7, 34, 2, 2, 722, 723, 5, 128, 65, 2, 723, 724, 7,
	35, 2, 2, 724, 733, 3, 2, 2, 2, 725, 726, 7, 37, 2, 2, 726, 727, 7, 34,
	2, 2, 727, 728, 5, 128, 65, 2, 728, 729, 7, 6, 2, 2, 729, 730, 5, 128,
	65, 2, 730, 731, 7, 35, 2, 2, 731, 733, 3, 2, 2, 2, 732, 715, 3, 2, 2,
	2, 732, 720, 3, 2, 2, 2, 732, 725, 3, 2, 2, 2, 733, 131, 3, 2, 2, 2, 734,
	735, 9, 6, 2, 2, 735, 133, 3, 2, 2, 2, 76, 138, 143, 149, 165, 171, 176,
	182, 187, 192, 207, 217, 223, 243, 248, 257, 261, 264, 275, 277, 282, 293,
	296, 304, 315, 326, 333, 346, 353, 360, 374, 389, 397, 410, 417, 432, 435,
	443, 447, 454, 459, 468, 471, 478, 488, 496, 500, 504, 516, 525, 536, 544,
	552, 560, 564, 569, 579, 591, 596, 613, 618, 624, 629, 642, 645, 655, 658,
	668, 671, 682, 685, 694, 706, 711, 732,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'('", "')'", "','", "'='", "'+'", "'-'", "'['", "']'", "'?'",
	"'.'", "':'", "'{'", "'}'", "'ascii'", "'bigint'", "'blob'", "'boolean'",
	"'counter'", "'decimal'", "'double'", "'float'", "'inet'", "'int'", "'text'",
	"'timestamp'", "'timeuuid'", "'uuid'", "'varchar'", "'varint'", "'list'",
	"'<'", "'>'", "'set'", "'map'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"K_ADD", "K_ALTER", "K_AND", "K_APPLY", "K_BATCH", "K_BEGIN", "K_CLUSTERING",
	"K_COLUMNFAMILY", "K_COMPACT", "K_COUNTER", "K_CREATE", "K_CUSTOM", "K_DELETE",
	"K_DROP", "K_EXISTS", "K_FALSE", "K_FROM", "K_IF", "K_IN", "K_INDEX", "K_INSERT",
	"K_INTO", "K_KEY", "K_KEYSPACE", "K_NOT", "K_ON", "K_OPTIONS", "K_ORDER",
	"K_PRIMARY", "K_SELECT", "K_SET", "K_STATIC", "K_STORAGE", "K_TABLE", "K_TIMESTAMP",
	"K_TRUE", "K_TRUNCATE", "K_TTL", "K_TYPE", "K_UNLOGGED", "K_UPDATE", "K_USE",
	"K_USING", "K_VALUES", "K_WHERE", "K_WITH", "IDENTIFIER", "STRING", "INTEGER",
	"FLOAT", "UUID", "BLOB", "SINGLE_LINE_COMMENT", "MULTILINE_COMMENT", "WS",
}

var ruleNames = []string{
	"statements", "statement", "dml_statements", "dml_statement", "create_keyspace_stmt",
	"alter_keyspace_stmt", "drop_keyspace_stmt", "use_stmt", "create_table_stmt",
	"alter_table_stmt", "alter_table_instruction", "drop_table_stmt", "truncate_table_stmt",
	"create_index_stmt", "drop_index_stmt", "insert_stmt", "column_names",
	"column_values", "upsert_options", "upsert_option", "index_name", "index_class",
	"index_options", "update_stmt", "update_assignments", "update_assignment",
	"update_conditions", "update_condition", "where_clause", "relation", "delete_stmt",
	"delete_conditions", "delete_condition", "delete_selections", "delete_selection",
	"batch_stmt", "batch_options", "batch_option", "table_name", "column_name",
	"table_options", "table_option", "column_definitions", "column_definition",
	"column_type", "primary_key", "partition_key", "keyspace_name", "if_not_exists",
	"if_exists", "constant", "variable", "term", "collection", "map", "set",
	"list", "function", "properties", "property", "property_name", "property_value",
	"data_type", "native_type", "collection_type", "bool",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CQL3Parser struct {
	*antlr.BaseParser
}

func NewCQL3Parser(input antlr.TokenStream) *CQL3Parser {
	this := new(CQL3Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CQL3.g4"

	return this
}

// CQL3Parser tokens.
const (
	CQL3ParserEOF                 = antlr.TokenEOF
	CQL3ParserT__0                = 1
	CQL3ParserT__1                = 2
	CQL3ParserT__2                = 3
	CQL3ParserT__3                = 4
	CQL3ParserT__4                = 5
	CQL3ParserT__5                = 6
	CQL3ParserT__6                = 7
	CQL3ParserT__7                = 8
	CQL3ParserT__8                = 9
	CQL3ParserT__9                = 10
	CQL3ParserT__10               = 11
	CQL3ParserT__11               = 12
	CQL3ParserT__12               = 13
	CQL3ParserT__13               = 14
	CQL3ParserT__14               = 15
	CQL3ParserT__15               = 16
	CQL3ParserT__16               = 17
	CQL3ParserT__17               = 18
	CQL3ParserT__18               = 19
	CQL3ParserT__19               = 20
	CQL3ParserT__20               = 21
	CQL3ParserT__21               = 22
	CQL3ParserT__22               = 23
	CQL3ParserT__23               = 24
	CQL3ParserT__24               = 25
	CQL3ParserT__25               = 26
	CQL3ParserT__26               = 27
	CQL3ParserT__27               = 28
	CQL3ParserT__28               = 29
	CQL3ParserT__29               = 30
	CQL3ParserT__30               = 31
	CQL3ParserT__31               = 32
	CQL3ParserT__32               = 33
	CQL3ParserT__33               = 34
	CQL3ParserT__34               = 35
	CQL3ParserK_ADD               = 36
	CQL3ParserK_ALTER             = 37
	CQL3ParserK_AND               = 38
	CQL3ParserK_APPLY             = 39
	CQL3ParserK_BATCH             = 40
	CQL3ParserK_BEGIN             = 41
	CQL3ParserK_CLUSTERING        = 42
	CQL3ParserK_COLUMNFAMILY      = 43
	CQL3ParserK_COMPACT           = 44
	CQL3ParserK_COUNTER           = 45
	CQL3ParserK_CREATE            = 46
	CQL3ParserK_CUSTOM            = 47
	CQL3ParserK_DELETE            = 48
	CQL3ParserK_DROP              = 49
	CQL3ParserK_EXISTS            = 50
	CQL3ParserK_FALSE             = 51
	CQL3ParserK_FROM              = 52
	CQL3ParserK_IF                = 53
	CQL3ParserK_IN                = 54
	CQL3ParserK_INDEX             = 55
	CQL3ParserK_INSERT            = 56
	CQL3ParserK_INTO              = 57
	CQL3ParserK_KEY               = 58
	CQL3ParserK_KEYSPACE          = 59
	CQL3ParserK_NOT               = 60
	CQL3ParserK_ON                = 61
	CQL3ParserK_OPTIONS           = 62
	CQL3ParserK_ORDER             = 63
	CQL3ParserK_PRIMARY           = 64
	CQL3ParserK_SELECT            = 65
	CQL3ParserK_SET               = 66
	CQL3ParserK_STATIC            = 67
	CQL3ParserK_STORAGE           = 68
	CQL3ParserK_TABLE             = 69
	CQL3ParserK_TIMESTAMP         = 70
	CQL3ParserK_TRUE              = 71
	CQL3ParserK_TRUNCATE          = 72
	CQL3ParserK_TTL               = 73
	CQL3ParserK_TYPE              = 74
	CQL3ParserK_UNLOGGED          = 75
	CQL3ParserK_UPDATE            = 76
	CQL3ParserK_USE               = 77
	CQL3ParserK_USING             = 78
	CQL3ParserK_VALUES            = 79
	CQL3ParserK_WHERE             = 80
	CQL3ParserK_WITH              = 81
	CQL3ParserIDENTIFIER          = 82
	CQL3ParserSTRING              = 83
	CQL3ParserINTEGER             = 84
	CQL3ParserFLOAT               = 85
	CQL3ParserUUID                = 86
	CQL3ParserBLOB                = 87
	CQL3ParserSINGLE_LINE_COMMENT = 88
	CQL3ParserMULTILINE_COMMENT   = 89
	CQL3ParserWS                  = 90
)

// CQL3Parser rules.
const (
	CQL3ParserRULE_statements              = 0
	CQL3ParserRULE_statement               = 1
	CQL3ParserRULE_dml_statements          = 2
	CQL3ParserRULE_dml_statement           = 3
	CQL3ParserRULE_create_keyspace_stmt    = 4
	CQL3ParserRULE_alter_keyspace_stmt     = 5
	CQL3ParserRULE_drop_keyspace_stmt      = 6
	CQL3ParserRULE_use_stmt                = 7
	CQL3ParserRULE_create_table_stmt       = 8
	CQL3ParserRULE_alter_table_stmt        = 9
	CQL3ParserRULE_alter_table_instruction = 10
	CQL3ParserRULE_drop_table_stmt         = 11
	CQL3ParserRULE_truncate_table_stmt     = 12
	CQL3ParserRULE_create_index_stmt       = 13
	CQL3ParserRULE_drop_index_stmt         = 14
	CQL3ParserRULE_insert_stmt             = 15
	CQL3ParserRULE_column_names            = 16
	CQL3ParserRULE_column_values           = 17
	CQL3ParserRULE_upsert_options          = 18
	CQL3ParserRULE_upsert_option           = 19
	CQL3ParserRULE_index_name              = 20
	CQL3ParserRULE_index_class             = 21
	CQL3ParserRULE_index_options           = 22
	CQL3ParserRULE_update_stmt             = 23
	CQL3ParserRULE_update_assignments      = 24
	CQL3ParserRULE_update_assignment       = 25
	CQL3ParserRULE_update_conditions       = 26
	CQL3ParserRULE_update_condition        = 27
	CQL3ParserRULE_where_clause            = 28
	CQL3ParserRULE_relation                = 29
	CQL3ParserRULE_delete_stmt             = 30
	CQL3ParserRULE_delete_conditions       = 31
	CQL3ParserRULE_delete_condition        = 32
	CQL3ParserRULE_delete_selections       = 33
	CQL3ParserRULE_delete_selection        = 34
	CQL3ParserRULE_batch_stmt              = 35
	CQL3ParserRULE_batch_options           = 36
	CQL3ParserRULE_batch_option            = 37
	CQL3ParserRULE_table_name              = 38
	CQL3ParserRULE_column_name             = 39
	CQL3ParserRULE_table_options           = 40
	CQL3ParserRULE_table_option            = 41
	CQL3ParserRULE_column_definitions      = 42
	CQL3ParserRULE_column_definition       = 43
	CQL3ParserRULE_column_type             = 44
	CQL3ParserRULE_primary_key             = 45
	CQL3ParserRULE_partition_key           = 46
	CQL3ParserRULE_keyspace_name           = 47
	CQL3ParserRULE_if_not_exists           = 48
	CQL3ParserRULE_if_exists               = 49
	CQL3ParserRULE_constant                = 50
	CQL3ParserRULE_variable                = 51
	CQL3ParserRULE_term                    = 52
	CQL3ParserRULE_collection              = 53
	CQL3ParserRULE_map                     = 54
	CQL3ParserRULE_set                     = 55
	CQL3ParserRULE_list                    = 56
	CQL3ParserRULE_function                = 57
	CQL3ParserRULE_properties              = 58
	CQL3ParserRULE_property                = 59
	CQL3ParserRULE_property_name           = 60
	CQL3ParserRULE_property_value          = 61
	CQL3ParserRULE_data_type               = 62
	CQL3ParserRULE_native_type             = 63
	CQL3ParserRULE_collection_type         = 64
	CQL3ParserRULE_bool                    = 65
)

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *CQL3Parser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CQL3ParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Statement()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == CQL3ParserT__0 {
				{
					p.SetState(133)
					p.Match(CQL3ParserT__0)
				}

				p.SetState(136)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(138)
				p.Statement()
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQL3ParserT__0 {
		{
			p.SetState(144)
			p.Match(CQL3ParserT__0)
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Drop_keyspace_stmt() IDrop_keyspace_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_keyspace_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_keyspace_stmtContext)
}

func (s *StatementContext) Create_keyspace_stmt() ICreate_keyspace_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_keyspace_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_keyspace_stmtContext)
}

func (s *StatementContext) Alter_keyspace_stmt() IAlter_keyspace_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlter_keyspace_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlter_keyspace_stmtContext)
}

func (s *StatementContext) Use_stmt() IUse_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUse_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUse_stmtContext)
}

func (s *StatementContext) Create_table_stmt() ICreate_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_table_stmtContext)
}

func (s *StatementContext) Alter_table_stmt() IAlter_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlter_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlter_table_stmtContext)
}

func (s *StatementContext) Drop_table_stmt() IDrop_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_table_stmtContext)
}

func (s *StatementContext) Truncate_table_stmt() ITruncate_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITruncate_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITruncate_table_stmtContext)
}

func (s *StatementContext) Create_index_stmt() ICreate_index_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_index_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_index_stmtContext)
}

func (s *StatementContext) Drop_index_stmt() IDrop_index_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_index_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_index_stmtContext)
}

func (s *StatementContext) Insert_stmt() IInsert_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert_stmtContext)
}

func (s *StatementContext) Update_stmt() IUpdate_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_stmtContext)
}

func (s *StatementContext) Delete_stmt() IDelete_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelete_stmtContext)
}

func (s *StatementContext) Batch_stmt() IBatch_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBatch_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBatch_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *CQL3Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CQL3ParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Drop_keyspace_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Create_keyspace_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Alter_keyspace_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Use_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Create_table_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(154)
			p.Alter_table_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Drop_table_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(156)
			p.Truncate_table_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(157)
			p.Create_index_stmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(158)
			p.Drop_index_stmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(159)
			p.Insert_stmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(160)
			p.Update_stmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(161)
			p.Delete_stmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(162)
			p.Batch_stmt()
		}

	}

	return localctx
}

// IDml_statementsContext is an interface to support dynamic dispatch.
type IDml_statementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDml_statementsContext differentiates from other interfaces.
	IsDml_statementsContext()
}

type Dml_statementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDml_statementsContext() *Dml_statementsContext {
	var p = new(Dml_statementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_dml_statements
	return p
}

func (*Dml_statementsContext) IsDml_statementsContext() {}

func NewDml_statementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dml_statementsContext {
	var p = new(Dml_statementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_dml_statements

	return p
}

func (s *Dml_statementsContext) GetParser() antlr.Parser { return s.parser }

func (s *Dml_statementsContext) AllDml_statement() []IDml_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDml_statementContext)(nil)).Elem())
	var tst = make([]IDml_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDml_statementContext)
		}
	}

	return tst
}

func (s *Dml_statementsContext) Dml_statement(i int) IDml_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDml_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDml_statementContext)
}

func (s *Dml_statementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dml_statementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dml_statementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDml_statements(s)
	}
}

func (s *Dml_statementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDml_statements(s)
	}
}

func (p *CQL3Parser) Dml_statements() (localctx IDml_statementsContext) {
	localctx = NewDml_statementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CQL3ParserRULE_dml_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Dml_statement()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == CQL3ParserT__0 {
				{
					p.SetState(166)
					p.Match(CQL3ParserT__0)
				}

				p.SetState(169)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(171)
				p.Dml_statement()
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQL3ParserT__0 {
		{
			p.SetState(177)
			p.Match(CQL3ParserT__0)
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDml_statementContext is an interface to support dynamic dispatch.
type IDml_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDml_statementContext differentiates from other interfaces.
	IsDml_statementContext()
}

type Dml_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDml_statementContext() *Dml_statementContext {
	var p = new(Dml_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_dml_statement
	return p
}

func (*Dml_statementContext) IsDml_statementContext() {}

func NewDml_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dml_statementContext {
	var p = new(Dml_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_dml_statement

	return p
}

func (s *Dml_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Dml_statementContext) Insert_stmt() IInsert_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert_stmtContext)
}

func (s *Dml_statementContext) Update_stmt() IUpdate_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_stmtContext)
}

func (s *Dml_statementContext) Delete_stmt() IDelete_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelete_stmtContext)
}

func (s *Dml_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dml_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dml_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDml_statement(s)
	}
}

func (s *Dml_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDml_statement(s)
	}
}

func (p *CQL3Parser) Dml_statement() (localctx IDml_statementContext) {
	localctx = NewDml_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CQL3ParserRULE_dml_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserK_INSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Insert_stmt()
		}

	case CQL3ParserK_UPDATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Update_stmt()
		}

	case CQL3ParserK_DELETE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Delete_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICreate_keyspace_stmtContext is an interface to support dynamic dispatch.
type ICreate_keyspace_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_keyspace_stmtContext differentiates from other interfaces.
	IsCreate_keyspace_stmtContext()
}

type Create_keyspace_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_keyspace_stmtContext() *Create_keyspace_stmtContext {
	var p = new(Create_keyspace_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_create_keyspace_stmt
	return p
}

func (*Create_keyspace_stmtContext) IsCreate_keyspace_stmtContext() {}

func NewCreate_keyspace_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_keyspace_stmtContext {
	var p = new(Create_keyspace_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_create_keyspace_stmt

	return p
}

func (s *Create_keyspace_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_keyspace_stmtContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_CREATE, 0)
}

func (s *Create_keyspace_stmtContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_KEYSPACE, 0)
}

func (s *Create_keyspace_stmtContext) Keyspace_name() IKeyspace_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyspace_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyspace_nameContext)
}

func (s *Create_keyspace_stmtContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WITH, 0)
}

func (s *Create_keyspace_stmtContext) Properties() IPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *Create_keyspace_stmtContext) If_not_exists() IIf_not_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_not_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_not_existsContext)
}

func (s *Create_keyspace_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_keyspace_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_keyspace_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterCreate_keyspace_stmt(s)
	}
}

func (s *Create_keyspace_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitCreate_keyspace_stmt(s)
	}
}

func (p *CQL3Parser) Create_keyspace_stmt() (localctx ICreate_keyspace_stmtContext) {
	localctx = NewCreate_keyspace_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CQL3ParserRULE_create_keyspace_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(CQL3ParserK_CREATE)
	}
	{
		p.SetState(188)
		p.Match(CQL3ParserK_KEYSPACE)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(189)
			p.If_not_exists()
		}

	}
	{
		p.SetState(192)
		p.Keyspace_name()
	}
	{
		p.SetState(193)
		p.Match(CQL3ParserK_WITH)
	}
	{
		p.SetState(194)
		p.Properties()
	}

	return localctx
}

// IAlter_keyspace_stmtContext is an interface to support dynamic dispatch.
type IAlter_keyspace_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlter_keyspace_stmtContext differentiates from other interfaces.
	IsAlter_keyspace_stmtContext()
}

type Alter_keyspace_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlter_keyspace_stmtContext() *Alter_keyspace_stmtContext {
	var p = new(Alter_keyspace_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_alter_keyspace_stmt
	return p
}

func (*Alter_keyspace_stmtContext) IsAlter_keyspace_stmtContext() {}

func NewAlter_keyspace_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alter_keyspace_stmtContext {
	var p = new(Alter_keyspace_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_alter_keyspace_stmt

	return p
}

func (s *Alter_keyspace_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Alter_keyspace_stmtContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_ALTER, 0)
}

func (s *Alter_keyspace_stmtContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_KEYSPACE, 0)
}

func (s *Alter_keyspace_stmtContext) Keyspace_name() IKeyspace_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyspace_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyspace_nameContext)
}

func (s *Alter_keyspace_stmtContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WITH, 0)
}

func (s *Alter_keyspace_stmtContext) Properties() IPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *Alter_keyspace_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alter_keyspace_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alter_keyspace_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterAlter_keyspace_stmt(s)
	}
}

func (s *Alter_keyspace_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitAlter_keyspace_stmt(s)
	}
}

func (p *CQL3Parser) Alter_keyspace_stmt() (localctx IAlter_keyspace_stmtContext) {
	localctx = NewAlter_keyspace_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CQL3ParserRULE_alter_keyspace_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(CQL3ParserK_ALTER)
	}
	{
		p.SetState(197)
		p.Match(CQL3ParserK_KEYSPACE)
	}
	{
		p.SetState(198)
		p.Keyspace_name()
	}
	{
		p.SetState(199)
		p.Match(CQL3ParserK_WITH)
	}
	{
		p.SetState(200)
		p.Properties()
	}

	return localctx
}

// IDrop_keyspace_stmtContext is an interface to support dynamic dispatch.
type IDrop_keyspace_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDrop_keyspace_stmtContext differentiates from other interfaces.
	IsDrop_keyspace_stmtContext()
}

type Drop_keyspace_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrop_keyspace_stmtContext() *Drop_keyspace_stmtContext {
	var p = new(Drop_keyspace_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_drop_keyspace_stmt
	return p
}

func (*Drop_keyspace_stmtContext) IsDrop_keyspace_stmtContext() {}

func NewDrop_keyspace_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_keyspace_stmtContext {
	var p = new(Drop_keyspace_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_drop_keyspace_stmt

	return p
}

func (s *Drop_keyspace_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_keyspace_stmtContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_DROP, 0)
}

func (s *Drop_keyspace_stmtContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_KEYSPACE, 0)
}

func (s *Drop_keyspace_stmtContext) Keyspace_name() IKeyspace_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyspace_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyspace_nameContext)
}

func (s *Drop_keyspace_stmtContext) If_exists() IIf_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_existsContext)
}

func (s *Drop_keyspace_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_keyspace_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_keyspace_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDrop_keyspace_stmt(s)
	}
}

func (s *Drop_keyspace_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDrop_keyspace_stmt(s)
	}
}

func (p *CQL3Parser) Drop_keyspace_stmt() (localctx IDrop_keyspace_stmtContext) {
	localctx = NewDrop_keyspace_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CQL3ParserRULE_drop_keyspace_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(CQL3ParserK_DROP)
	}
	{
		p.SetState(203)
		p.Match(CQL3ParserK_KEYSPACE)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(204)
			p.If_exists()
		}

	}
	{
		p.SetState(207)
		p.Keyspace_name()
	}

	return localctx
}

// IUse_stmtContext is an interface to support dynamic dispatch.
type IUse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUse_stmtContext differentiates from other interfaces.
	IsUse_stmtContext()
}

type Use_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUse_stmtContext() *Use_stmtContext {
	var p = new(Use_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_use_stmt
	return p
}

func (*Use_stmtContext) IsUse_stmtContext() {}

func NewUse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Use_stmtContext {
	var p = new(Use_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_use_stmt

	return p
}

func (s *Use_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Use_stmtContext) K_USE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_USE, 0)
}

func (s *Use_stmtContext) Keyspace_name() IKeyspace_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyspace_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyspace_nameContext)
}

func (s *Use_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Use_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Use_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUse_stmt(s)
	}
}

func (s *Use_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUse_stmt(s)
	}
}

func (p *CQL3Parser) Use_stmt() (localctx IUse_stmtContext) {
	localctx = NewUse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQL3ParserRULE_use_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(CQL3ParserK_USE)
	}
	{
		p.SetState(210)
		p.Keyspace_name()
	}

	return localctx
}

// ICreate_table_stmtContext is an interface to support dynamic dispatch.
type ICreate_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_table_stmtContext differentiates from other interfaces.
	IsCreate_table_stmtContext()
}

type Create_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_table_stmtContext() *Create_table_stmtContext {
	var p = new(Create_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_create_table_stmt
	return p
}

func (*Create_table_stmtContext) IsCreate_table_stmtContext() {}

func NewCreate_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_table_stmtContext {
	var p = new(Create_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_create_table_stmt

	return p
}

func (s *Create_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_table_stmtContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_CREATE, 0)
}

func (s *Create_table_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Create_table_stmtContext) Column_definitions() IColumn_definitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_definitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_definitionsContext)
}

func (s *Create_table_stmtContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TABLE, 0)
}

func (s *Create_table_stmtContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_COLUMNFAMILY, 0)
}

func (s *Create_table_stmtContext) If_not_exists() IIf_not_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_not_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_not_existsContext)
}

func (s *Create_table_stmtContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WITH, 0)
}

func (s *Create_table_stmtContext) Table_options() ITable_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_optionsContext)
}

func (s *Create_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterCreate_table_stmt(s)
	}
}

func (s *Create_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitCreate_table_stmt(s)
	}
}

func (p *CQL3Parser) Create_table_stmt() (localctx ICreate_table_stmtContext) {
	localctx = NewCreate_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CQL3ParserRULE_create_table_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(CQL3ParserK_CREATE)
	}
	p.SetState(213)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQL3ParserK_COLUMNFAMILY || _la == CQL3ParserK_TABLE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(214)
			p.If_not_exists()
		}

	}
	{
		p.SetState(217)
		p.Table_name()
	}
	{
		p.SetState(218)
		p.Column_definitions()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_WITH {
		{
			p.SetState(219)
			p.Match(CQL3ParserK_WITH)
		}
		{
			p.SetState(220)
			p.Table_options()
		}

	}

	return localctx
}

// IAlter_table_stmtContext is an interface to support dynamic dispatch.
type IAlter_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlter_table_stmtContext differentiates from other interfaces.
	IsAlter_table_stmtContext()
}

type Alter_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlter_table_stmtContext() *Alter_table_stmtContext {
	var p = new(Alter_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_alter_table_stmt
	return p
}

func (*Alter_table_stmtContext) IsAlter_table_stmtContext() {}

func NewAlter_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alter_table_stmtContext {
	var p = new(Alter_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_alter_table_stmt

	return p
}

func (s *Alter_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Alter_table_stmtContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_ALTER, 0)
}

func (s *Alter_table_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Alter_table_stmtContext) Alter_table_instruction() IAlter_table_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlter_table_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlter_table_instructionContext)
}

func (s *Alter_table_stmtContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TABLE, 0)
}

func (s *Alter_table_stmtContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_COLUMNFAMILY, 0)
}

func (s *Alter_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alter_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alter_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterAlter_table_stmt(s)
	}
}

func (s *Alter_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitAlter_table_stmt(s)
	}
}

func (p *CQL3Parser) Alter_table_stmt() (localctx IAlter_table_stmtContext) {
	localctx = NewAlter_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQL3ParserRULE_alter_table_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(CQL3ParserK_ALTER)
	}
	p.SetState(224)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQL3ParserK_COLUMNFAMILY || _la == CQL3ParserK_TABLE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(225)
		p.Table_name()
	}
	{
		p.SetState(226)
		p.Alter_table_instruction()
	}

	return localctx
}

// IAlter_table_instructionContext is an interface to support dynamic dispatch.
type IAlter_table_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlter_table_instructionContext differentiates from other interfaces.
	IsAlter_table_instructionContext()
}

type Alter_table_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlter_table_instructionContext() *Alter_table_instructionContext {
	var p = new(Alter_table_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_alter_table_instruction
	return p
}

func (*Alter_table_instructionContext) IsAlter_table_instructionContext() {}

func NewAlter_table_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alter_table_instructionContext {
	var p = new(Alter_table_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_alter_table_instruction

	return p
}

func (s *Alter_table_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Alter_table_instructionContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_ALTER, 0)
}

func (s *Alter_table_instructionContext) Column_name() IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Alter_table_instructionContext) K_TYPE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TYPE, 0)
}

func (s *Alter_table_instructionContext) Column_type() IColumn_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_typeContext)
}

func (s *Alter_table_instructionContext) K_ADD() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_ADD, 0)
}

func (s *Alter_table_instructionContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_DROP, 0)
}

func (s *Alter_table_instructionContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WITH, 0)
}

func (s *Alter_table_instructionContext) Table_options() ITable_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_optionsContext)
}

func (s *Alter_table_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alter_table_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alter_table_instructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterAlter_table_instruction(s)
	}
}

func (s *Alter_table_instructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitAlter_table_instruction(s)
	}
}

func (p *CQL3Parser) Alter_table_instruction() (localctx IAlter_table_instructionContext) {
	localctx = NewAlter_table_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CQL3ParserRULE_alter_table_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(241)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserK_ALTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(CQL3ParserK_ALTER)
		}
		{
			p.SetState(229)
			p.Column_name()
		}
		{
			p.SetState(230)
			p.Match(CQL3ParserK_TYPE)
		}
		{
			p.SetState(231)
			p.Column_type()
		}

	case CQL3ParserK_ADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Match(CQL3ParserK_ADD)
		}
		{
			p.SetState(234)
			p.Column_name()
		}
		{
			p.SetState(235)
			p.Column_type()
		}

	case CQL3ParserK_DROP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.Match(CQL3ParserK_DROP)
		}
		{
			p.SetState(238)
			p.Column_name()
		}

	case CQL3ParserK_WITH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(239)
			p.Match(CQL3ParserK_WITH)
		}
		{
			p.SetState(240)
			p.Table_options()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDrop_table_stmtContext is an interface to support dynamic dispatch.
type IDrop_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDrop_table_stmtContext differentiates from other interfaces.
	IsDrop_table_stmtContext()
}

type Drop_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrop_table_stmtContext() *Drop_table_stmtContext {
	var p = new(Drop_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_drop_table_stmt
	return p
}

func (*Drop_table_stmtContext) IsDrop_table_stmtContext() {}

func NewDrop_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_table_stmtContext {
	var p = new(Drop_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_drop_table_stmt

	return p
}

func (s *Drop_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_table_stmtContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_DROP, 0)
}

func (s *Drop_table_stmtContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TABLE, 0)
}

func (s *Drop_table_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Drop_table_stmtContext) If_exists() IIf_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_existsContext)
}

func (s *Drop_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDrop_table_stmt(s)
	}
}

func (s *Drop_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDrop_table_stmt(s)
	}
}

func (p *CQL3Parser) Drop_table_stmt() (localctx IDrop_table_stmtContext) {
	localctx = NewDrop_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CQL3ParserRULE_drop_table_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(CQL3ParserK_DROP)
	}
	{
		p.SetState(244)
		p.Match(CQL3ParserK_TABLE)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(245)
			p.If_exists()
		}

	}
	{
		p.SetState(248)
		p.Table_name()
	}

	return localctx
}

// ITruncate_table_stmtContext is an interface to support dynamic dispatch.
type ITruncate_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTruncate_table_stmtContext differentiates from other interfaces.
	IsTruncate_table_stmtContext()
}

type Truncate_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTruncate_table_stmtContext() *Truncate_table_stmtContext {
	var p = new(Truncate_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_truncate_table_stmt
	return p
}

func (*Truncate_table_stmtContext) IsTruncate_table_stmtContext() {}

func NewTruncate_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Truncate_table_stmtContext {
	var p = new(Truncate_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_truncate_table_stmt

	return p
}

func (s *Truncate_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Truncate_table_stmtContext) K_TRUNCATE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TRUNCATE, 0)
}

func (s *Truncate_table_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Truncate_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Truncate_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Truncate_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterTruncate_table_stmt(s)
	}
}

func (s *Truncate_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitTruncate_table_stmt(s)
	}
}

func (p *CQL3Parser) Truncate_table_stmt() (localctx ITruncate_table_stmtContext) {
	localctx = NewTruncate_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CQL3ParserRULE_truncate_table_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(CQL3ParserK_TRUNCATE)
	}
	{
		p.SetState(251)
		p.Table_name()
	}

	return localctx
}

// ICreate_index_stmtContext is an interface to support dynamic dispatch.
type ICreate_index_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_index_stmtContext differentiates from other interfaces.
	IsCreate_index_stmtContext()
}

type Create_index_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_index_stmtContext() *Create_index_stmtContext {
	var p = new(Create_index_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_create_index_stmt
	return p
}

func (*Create_index_stmtContext) IsCreate_index_stmtContext() {}

func NewCreate_index_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_index_stmtContext {
	var p = new(Create_index_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_create_index_stmt

	return p
}

func (s *Create_index_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_index_stmtContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_CREATE, 0)
}

func (s *Create_index_stmtContext) K_INDEX() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_INDEX, 0)
}

func (s *Create_index_stmtContext) K_ON() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_ON, 0)
}

func (s *Create_index_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Create_index_stmtContext) Column_name() IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Create_index_stmtContext) K_CUSTOM() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_CUSTOM, 0)
}

func (s *Create_index_stmtContext) If_not_exists() IIf_not_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_not_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_not_existsContext)
}

func (s *Create_index_stmtContext) Index_name() IIndex_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_nameContext)
}

func (s *Create_index_stmtContext) K_USING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_USING, 0)
}

func (s *Create_index_stmtContext) Index_class() IIndex_classContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_classContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_classContext)
}

func (s *Create_index_stmtContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WITH, 0)
}

func (s *Create_index_stmtContext) Index_options() IIndex_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_optionsContext)
}

func (s *Create_index_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_index_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_index_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterCreate_index_stmt(s)
	}
}

func (s *Create_index_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitCreate_index_stmt(s)
	}
}

func (p *CQL3Parser) Create_index_stmt() (localctx ICreate_index_stmtContext) {
	localctx = NewCreate_index_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CQL3ParserRULE_create_index_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(CQL3ParserK_CREATE)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_CUSTOM {
		{
			p.SetState(254)
			p.Match(CQL3ParserK_CUSTOM)
		}

	}
	{
		p.SetState(257)
		p.Match(CQL3ParserK_INDEX)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(258)
			p.If_not_exists()
		}

	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserIDENTIFIER {
		{
			p.SetState(261)
			p.Index_name()
		}

	}
	{
		p.SetState(264)
		p.Match(CQL3ParserK_ON)
	}
	{
		p.SetState(265)
		p.Table_name()
	}
	{
		p.SetState(266)
		p.Match(CQL3ParserT__1)
	}
	{
		p.SetState(267)
		p.Column_name()
	}
	{
		p.SetState(268)
		p.Match(CQL3ParserT__2)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_USING {
		{
			p.SetState(269)
			p.Match(CQL3ParserK_USING)
		}
		{
			p.SetState(270)
			p.Index_class()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CQL3ParserK_WITH {
			{
				p.SetState(271)
				p.Match(CQL3ParserK_WITH)
			}
			{
				p.SetState(272)
				p.Index_options()
			}

		}

	}

	return localctx
}

// IDrop_index_stmtContext is an interface to support dynamic dispatch.
type IDrop_index_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDrop_index_stmtContext differentiates from other interfaces.
	IsDrop_index_stmtContext()
}

type Drop_index_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrop_index_stmtContext() *Drop_index_stmtContext {
	var p = new(Drop_index_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_drop_index_stmt
	return p
}

func (*Drop_index_stmtContext) IsDrop_index_stmtContext() {}

func NewDrop_index_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_index_stmtContext {
	var p = new(Drop_index_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_drop_index_stmt

	return p
}

func (s *Drop_index_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_index_stmtContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_DROP, 0)
}

func (s *Drop_index_stmtContext) K_INDEX() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_INDEX, 0)
}

func (s *Drop_index_stmtContext) Index_name() IIndex_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_nameContext)
}

func (s *Drop_index_stmtContext) If_exists() IIf_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_existsContext)
}

func (s *Drop_index_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_index_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_index_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDrop_index_stmt(s)
	}
}

func (s *Drop_index_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDrop_index_stmt(s)
	}
}

func (p *CQL3Parser) Drop_index_stmt() (localctx IDrop_index_stmtContext) {
	localctx = NewDrop_index_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CQL3ParserRULE_drop_index_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(CQL3ParserK_DROP)
	}
	{
		p.SetState(278)
		p.Match(CQL3ParserK_INDEX)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(279)
			p.If_exists()
		}

	}
	{
		p.SetState(282)
		p.Index_name()
	}

	return localctx
}

// IInsert_stmtContext is an interface to support dynamic dispatch.
type IInsert_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsert_stmtContext differentiates from other interfaces.
	IsInsert_stmtContext()
}

type Insert_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_stmtContext() *Insert_stmtContext {
	var p = new(Insert_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_insert_stmt
	return p
}

func (*Insert_stmtContext) IsInsert_stmtContext() {}

func NewInsert_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_stmtContext {
	var p = new(Insert_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_insert_stmt

	return p
}

func (s *Insert_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_stmtContext) K_INSERT() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_INSERT, 0)
}

func (s *Insert_stmtContext) K_INTO() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_INTO, 0)
}

func (s *Insert_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Insert_stmtContext) Column_names() IColumn_namesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_namesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_namesContext)
}

func (s *Insert_stmtContext) K_VALUES() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_VALUES, 0)
}

func (s *Insert_stmtContext) Column_values() IColumn_valuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_valuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_valuesContext)
}

func (s *Insert_stmtContext) If_not_exists() IIf_not_existsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_not_existsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_not_existsContext)
}

func (s *Insert_stmtContext) Upsert_options() IUpsert_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpsert_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpsert_optionsContext)
}

func (s *Insert_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterInsert_stmt(s)
	}
}

func (s *Insert_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitInsert_stmt(s)
	}
}

func (p *CQL3Parser) Insert_stmt() (localctx IInsert_stmtContext) {
	localctx = NewInsert_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CQL3ParserRULE_insert_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(CQL3ParserK_INSERT)
	}
	{
		p.SetState(285)
		p.Match(CQL3ParserK_INTO)
	}
	{
		p.SetState(286)
		p.Table_name()
	}
	{
		p.SetState(287)
		p.Column_names()
	}
	{
		p.SetState(288)
		p.Match(CQL3ParserK_VALUES)
	}
	{
		p.SetState(289)
		p.Column_values()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(290)
			p.If_not_exists()
		}

	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_USING {
		{
			p.SetState(293)
			p.Upsert_options()
		}

	}

	return localctx
}

// IColumn_namesContext is an interface to support dynamic dispatch.
type IColumn_namesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_namesContext differentiates from other interfaces.
	IsColumn_namesContext()
}

type Column_namesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_namesContext() *Column_namesContext {
	var p = new(Column_namesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_column_names
	return p
}

func (*Column_namesContext) IsColumn_namesContext() {}

func NewColumn_namesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_namesContext {
	var p = new(Column_namesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_column_names

	return p
}

func (s *Column_namesContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_namesContext) AllColumn_name() []IColumn_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem())
	var tst = make([]IColumn_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_nameContext)
		}
	}

	return tst
}

func (s *Column_namesContext) Column_name(i int) IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_namesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_namesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_namesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterColumn_names(s)
	}
}

func (s *Column_namesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitColumn_names(s)
	}
}

func (p *CQL3Parser) Column_names() (localctx IColumn_namesContext) {
	localctx = NewColumn_namesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQL3ParserRULE_column_names)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(CQL3ParserT__1)
	}
	{
		p.SetState(297)
		p.Column_name()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserT__3 {
		{
			p.SetState(298)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(299)
			p.Column_name()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(305)
		p.Match(CQL3ParserT__2)
	}

	return localctx
}

// IColumn_valuesContext is an interface to support dynamic dispatch.
type IColumn_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_valuesContext differentiates from other interfaces.
	IsColumn_valuesContext()
}

type Column_valuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_valuesContext() *Column_valuesContext {
	var p = new(Column_valuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_column_values
	return p
}

func (*Column_valuesContext) IsColumn_valuesContext() {}

func NewColumn_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_valuesContext {
	var p = new(Column_valuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_column_values

	return p
}

func (s *Column_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_valuesContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Column_valuesContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Column_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterColumn_values(s)
	}
}

func (s *Column_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitColumn_values(s)
	}
}

func (p *CQL3Parser) Column_values() (localctx IColumn_valuesContext) {
	localctx = NewColumn_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CQL3ParserRULE_column_values)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(CQL3ParserT__1)
	}
	{
		p.SetState(308)
		p.Term()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserT__3 {
		{
			p.SetState(309)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(310)
			p.Term()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(316)
		p.Match(CQL3ParserT__2)
	}

	return localctx
}

// IUpsert_optionsContext is an interface to support dynamic dispatch.
type IUpsert_optionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpsert_optionsContext differentiates from other interfaces.
	IsUpsert_optionsContext()
}

type Upsert_optionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpsert_optionsContext() *Upsert_optionsContext {
	var p = new(Upsert_optionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_upsert_options
	return p
}

func (*Upsert_optionsContext) IsUpsert_optionsContext() {}

func NewUpsert_optionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upsert_optionsContext {
	var p = new(Upsert_optionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_upsert_options

	return p
}

func (s *Upsert_optionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Upsert_optionsContext) K_USING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_USING, 0)
}

func (s *Upsert_optionsContext) AllUpsert_option() []IUpsert_optionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpsert_optionContext)(nil)).Elem())
	var tst = make([]IUpsert_optionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpsert_optionContext)
		}
	}

	return tst
}

func (s *Upsert_optionsContext) Upsert_option(i int) IUpsert_optionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpsert_optionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpsert_optionContext)
}

func (s *Upsert_optionsContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *Upsert_optionsContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *Upsert_optionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upsert_optionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upsert_optionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpsert_options(s)
	}
}

func (s *Upsert_optionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpsert_options(s)
	}
}

func (p *CQL3Parser) Upsert_options() (localctx IUpsert_optionsContext) {
	localctx = NewUpsert_optionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CQL3ParserRULE_upsert_options)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(CQL3ParserK_USING)
	}
	{
		p.SetState(319)
		p.Upsert_option()
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserK_AND {
		{
			p.SetState(320)
			p.Match(CQL3ParserK_AND)
		}
		{
			p.SetState(321)
			p.Upsert_option()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUpsert_optionContext is an interface to support dynamic dispatch.
type IUpsert_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpsert_optionContext differentiates from other interfaces.
	IsUpsert_optionContext()
}

type Upsert_optionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpsert_optionContext() *Upsert_optionContext {
	var p = new(Upsert_optionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_upsert_option
	return p
}

func (*Upsert_optionContext) IsUpsert_optionContext() {}

func NewUpsert_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upsert_optionContext {
	var p = new(Upsert_optionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_upsert_option

	return p
}

func (s *Upsert_optionContext) GetParser() antlr.Parser { return s.parser }

func (s *Upsert_optionContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TIMESTAMP, 0)
}

func (s *Upsert_optionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserINTEGER, 0)
}

func (s *Upsert_optionContext) K_TTL() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TTL, 0)
}

func (s *Upsert_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upsert_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upsert_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpsert_option(s)
	}
}

func (s *Upsert_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpsert_option(s)
	}
}

func (p *CQL3Parser) Upsert_option() (localctx IUpsert_optionContext) {
	localctx = NewUpsert_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CQL3ParserRULE_upsert_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(331)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserK_TIMESTAMP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(CQL3ParserK_TIMESTAMP)
		}
		{
			p.SetState(328)
			p.Match(CQL3ParserINTEGER)
		}

	case CQL3ParserK_TTL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(CQL3ParserK_TTL)
		}
		{
			p.SetState(330)
			p.Match(CQL3ParserINTEGER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndex_nameContext is an interface to support dynamic dispatch.
type IIndex_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndex_nameContext differentiates from other interfaces.
	IsIndex_nameContext()
}

type Index_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_nameContext() *Index_nameContext {
	var p = new(Index_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_index_name
	return p
}

func (*Index_nameContext) IsIndex_nameContext() {}

func NewIndex_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_nameContext {
	var p = new(Index_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_index_name

	return p
}

func (s *Index_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Index_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterIndex_name(s)
	}
}

func (s *Index_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitIndex_name(s)
	}
}

func (p *CQL3Parser) Index_name() (localctx IIndex_nameContext) {
	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CQL3ParserRULE_index_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(CQL3ParserIDENTIFIER)
	}

	return localctx
}

// IIndex_classContext is an interface to support dynamic dispatch.
type IIndex_classContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndex_classContext differentiates from other interfaces.
	IsIndex_classContext()
}

type Index_classContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_classContext() *Index_classContext {
	var p = new(Index_classContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_index_class
	return p
}

func (*Index_classContext) IsIndex_classContext() {}

func NewIndex_classContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_classContext {
	var p = new(Index_classContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_index_class

	return p
}

func (s *Index_classContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_classContext) STRING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserSTRING, 0)
}

func (s *Index_classContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_classContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_classContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterIndex_class(s)
	}
}

func (s *Index_classContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitIndex_class(s)
	}
}

func (p *CQL3Parser) Index_class() (localctx IIndex_classContext) {
	localctx = NewIndex_classContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CQL3ParserRULE_index_class)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(CQL3ParserSTRING)
	}

	return localctx
}

// IIndex_optionsContext is an interface to support dynamic dispatch.
type IIndex_optionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndex_optionsContext differentiates from other interfaces.
	IsIndex_optionsContext()
}

type Index_optionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_optionsContext() *Index_optionsContext {
	var p = new(Index_optionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_index_options
	return p
}

func (*Index_optionsContext) IsIndex_optionsContext() {}

func NewIndex_optionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_optionsContext {
	var p = new(Index_optionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_index_options

	return p
}

func (s *Index_optionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_optionsContext) K_OPTIONS() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_OPTIONS, 0)
}

func (s *Index_optionsContext) Map() IMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *Index_optionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_optionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_optionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterIndex_options(s)
	}
}

func (s *Index_optionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitIndex_options(s)
	}
}

func (p *CQL3Parser) Index_options() (localctx IIndex_optionsContext) {
	localctx = NewIndex_optionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CQL3ParserRULE_index_options)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(CQL3ParserK_OPTIONS)
	}
	{
		p.SetState(338)
		p.Match(CQL3ParserT__4)
	}
	{
		p.SetState(339)
		p.Map()
	}

	return localctx
}

// IUpdate_stmtContext is an interface to support dynamic dispatch.
type IUpdate_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_stmtContext differentiates from other interfaces.
	IsUpdate_stmtContext()
}

type Update_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_stmtContext() *Update_stmtContext {
	var p = new(Update_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_update_stmt
	return p
}

func (*Update_stmtContext) IsUpdate_stmtContext() {}

func NewUpdate_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_stmtContext {
	var p = new(Update_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_update_stmt

	return p
}

func (s *Update_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_stmtContext) K_UPDATE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_UPDATE, 0)
}

func (s *Update_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Update_stmtContext) K_SET() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_SET, 0)
}

func (s *Update_stmtContext) Update_assignments() IUpdate_assignmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_assignmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_assignmentsContext)
}

func (s *Update_stmtContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WHERE, 0)
}

func (s *Update_stmtContext) Where_clause() IWhere_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhere_clauseContext)
}

func (s *Update_stmtContext) Upsert_options() IUpsert_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpsert_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpsert_optionsContext)
}

func (s *Update_stmtContext) Update_conditions() IUpdate_conditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_conditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_conditionsContext)
}

func (s *Update_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpdate_stmt(s)
	}
}

func (s *Update_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpdate_stmt(s)
	}
}

func (p *CQL3Parser) Update_stmt() (localctx IUpdate_stmtContext) {
	localctx = NewUpdate_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CQL3ParserRULE_update_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(CQL3ParserK_UPDATE)
	}
	{
		p.SetState(342)
		p.Table_name()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_USING {
		{
			p.SetState(343)
			p.Upsert_options()
		}

	}
	{
		p.SetState(346)
		p.Match(CQL3ParserK_SET)
	}
	{
		p.SetState(347)
		p.Update_assignments()
	}
	{
		p.SetState(348)
		p.Match(CQL3ParserK_WHERE)
	}
	{
		p.SetState(349)
		p.Where_clause()
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(350)
			p.Update_conditions()
		}

	}

	return localctx
}

// IUpdate_assignmentsContext is an interface to support dynamic dispatch.
type IUpdate_assignmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_assignmentsContext differentiates from other interfaces.
	IsUpdate_assignmentsContext()
}

type Update_assignmentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_assignmentsContext() *Update_assignmentsContext {
	var p = new(Update_assignmentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_update_assignments
	return p
}

func (*Update_assignmentsContext) IsUpdate_assignmentsContext() {}

func NewUpdate_assignmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_assignmentsContext {
	var p = new(Update_assignmentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_update_assignments

	return p
}

func (s *Update_assignmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_assignmentsContext) AllUpdate_assignment() []IUpdate_assignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdate_assignmentContext)(nil)).Elem())
	var tst = make([]IUpdate_assignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdate_assignmentContext)
		}
	}

	return tst
}

func (s *Update_assignmentsContext) Update_assignment(i int) IUpdate_assignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_assignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdate_assignmentContext)
}

func (s *Update_assignmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_assignmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_assignmentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpdate_assignments(s)
	}
}

func (s *Update_assignmentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpdate_assignments(s)
	}
}

func (p *CQL3Parser) Update_assignments() (localctx IUpdate_assignmentsContext) {
	localctx = NewUpdate_assignmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CQL3ParserRULE_update_assignments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Update_assignment()
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserT__3 {
		{
			p.SetState(354)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(355)
			p.Update_assignment()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUpdate_assignmentContext is an interface to support dynamic dispatch.
type IUpdate_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_assignmentContext differentiates from other interfaces.
	IsUpdate_assignmentContext()
}

type Update_assignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_assignmentContext() *Update_assignmentContext {
	var p = new(Update_assignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_update_assignment
	return p
}

func (*Update_assignmentContext) IsUpdate_assignmentContext() {}

func NewUpdate_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_assignmentContext {
	var p = new(Update_assignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_update_assignment

	return p
}

func (s *Update_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_assignmentContext) AllColumn_name() []IColumn_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem())
	var tst = make([]IColumn_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_nameContext)
		}
	}

	return tst
}

func (s *Update_assignmentContext) Column_name(i int) IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Update_assignmentContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Update_assignmentContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Update_assignmentContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserINTEGER, 0)
}

func (s *Update_assignmentContext) Set() ISetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *Update_assignmentContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *Update_assignmentContext) Map() IMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *Update_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpdate_assignment(s)
	}
}

func (s *Update_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpdate_assignment(s)
	}
}

func (p *CQL3Parser) Update_assignment() (localctx IUpdate_assignmentContext) {
	localctx = NewUpdate_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CQL3ParserRULE_update_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.Column_name()
		}
		{
			p.SetState(362)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(363)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Column_name()
		}
		{
			p.SetState(366)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(367)
			p.Column_name()
		}
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CQL3ParserT__5 || _la == CQL3ParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CQL3ParserINTEGER:
			{
				p.SetState(369)
				p.Match(CQL3ParserINTEGER)
			}

		case CQL3ParserT__12:
			{
				p.SetState(370)
				p.Set()
			}

		case CQL3ParserT__7:
			{
				p.SetState(371)
				p.List()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(374)
			p.Column_name()
		}
		{
			p.SetState(375)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(376)
			p.Column_name()
		}
		{
			p.SetState(377)
			p.Match(CQL3ParserT__5)
		}
		{
			p.SetState(378)
			p.Map()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(380)
			p.Column_name()
		}
		{
			p.SetState(381)
			p.Match(CQL3ParserT__7)
		}
		{
			p.SetState(382)
			p.Term()
		}
		{
			p.SetState(383)
			p.Match(CQL3ParserT__8)
		}
		{
			p.SetState(384)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(385)
			p.Term()
		}

	}

	return localctx
}

// IUpdate_conditionsContext is an interface to support dynamic dispatch.
type IUpdate_conditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_conditionsContext differentiates from other interfaces.
	IsUpdate_conditionsContext()
}

type Update_conditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_conditionsContext() *Update_conditionsContext {
	var p = new(Update_conditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_update_conditions
	return p
}

func (*Update_conditionsContext) IsUpdate_conditionsContext() {}

func NewUpdate_conditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_conditionsContext {
	var p = new(Update_conditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_update_conditions

	return p
}

func (s *Update_conditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_conditionsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_IF, 0)
}

func (s *Update_conditionsContext) AllUpdate_condition() []IUpdate_conditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdate_conditionContext)(nil)).Elem())
	var tst = make([]IUpdate_conditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdate_conditionContext)
		}
	}

	return tst
}

func (s *Update_conditionsContext) Update_condition(i int) IUpdate_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_conditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdate_conditionContext)
}

func (s *Update_conditionsContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *Update_conditionsContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *Update_conditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_conditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_conditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpdate_conditions(s)
	}
}

func (s *Update_conditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpdate_conditions(s)
	}
}

func (p *CQL3Parser) Update_conditions() (localctx IUpdate_conditionsContext) {
	localctx = NewUpdate_conditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CQL3ParserRULE_update_conditions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(CQL3ParserK_IF)
	}
	{
		p.SetState(390)
		p.Update_condition()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserK_AND {
		{
			p.SetState(391)
			p.Match(CQL3ParserK_AND)
		}
		{
			p.SetState(392)
			p.Update_condition()
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUpdate_conditionContext is an interface to support dynamic dispatch.
type IUpdate_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_conditionContext differentiates from other interfaces.
	IsUpdate_conditionContext()
}

type Update_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_conditionContext() *Update_conditionContext {
	var p = new(Update_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_update_condition
	return p
}

func (*Update_conditionContext) IsUpdate_conditionContext() {}

func NewUpdate_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_conditionContext {
	var p = new(Update_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_update_condition

	return p
}

func (s *Update_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_conditionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Update_conditionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Update_conditionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Update_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterUpdate_condition(s)
	}
}

func (s *Update_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitUpdate_condition(s)
	}
}

func (p *CQL3Parser) Update_condition() (localctx IUpdate_conditionContext) {
	localctx = NewUpdate_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CQL3ParserRULE_update_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Match(CQL3ParserIDENTIFIER)
		}
		{
			p.SetState(399)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(400)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.Match(CQL3ParserIDENTIFIER)
		}
		{
			p.SetState(402)
			p.Match(CQL3ParserT__7)
		}
		{
			p.SetState(403)
			p.Term()
		}
		{
			p.SetState(404)
			p.Match(CQL3ParserT__8)
		}
		{
			p.SetState(405)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(406)
			p.Term()
		}

	}

	return localctx
}

// IWhere_clauseContext is an interface to support dynamic dispatch.
type IWhere_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhere_clauseContext differentiates from other interfaces.
	IsWhere_clauseContext()
}

type Where_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_clauseContext() *Where_clauseContext {
	var p = new(Where_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_where_clause
	return p
}

func (*Where_clauseContext) IsWhere_clauseContext() {}

func NewWhere_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_clauseContext {
	var p = new(Where_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_where_clause

	return p
}

func (s *Where_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_clauseContext) AllRelation() []IRelationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationContext)(nil)).Elem())
	var tst = make([]IRelationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationContext)
		}
	}

	return tst
}

func (s *Where_clauseContext) Relation(i int) IRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *Where_clauseContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *Where_clauseContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *Where_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Where_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterWhere_clause(s)
	}
}

func (s *Where_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitWhere_clause(s)
	}
}

func (p *CQL3Parser) Where_clause() (localctx IWhere_clauseContext) {
	localctx = NewWhere_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CQL3ParserRULE_where_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Relation()
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserK_AND {
		{
			p.SetState(411)
			p.Match(CQL3ParserK_AND)
		}
		{
			p.SetState(412)
			p.Relation()
		}

		p.SetState(417)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) Column_name() IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *RelationContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *RelationContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *RelationContext) K_IN() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_IN, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *CQL3Parser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CQL3ParserRULE_relation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Column_name()
		}
		{
			p.SetState(419)
			p.Match(CQL3ParserT__4)
		}
		{
			p.SetState(420)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.Column_name()
		}
		{
			p.SetState(423)
			p.Match(CQL3ParserK_IN)
		}
		{
			p.SetState(424)
			p.Match(CQL3ParserT__1)
		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQL3ParserT__7)|(1<<CQL3ParserT__9)|(1<<CQL3ParserT__11)|(1<<CQL3ParserT__12))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(CQL3ParserK_FALSE-51))|(1<<(CQL3ParserK_TRUE-51))|(1<<(CQL3ParserIDENTIFIER-51)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(CQL3ParserSTRING-83))|(1<<(CQL3ParserINTEGER-83))|(1<<(CQL3ParserFLOAT-83))|(1<<(CQL3ParserUUID-83))|(1<<(CQL3ParserBLOB-83)))) != 0) {
			{
				p.SetState(425)
				p.Term()
			}
			p.SetState(430)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == CQL3ParserT__3 {
				{
					p.SetState(426)
					p.Match(CQL3ParserT__3)
				}
				{
					p.SetState(427)
					p.Term()
				}

				p.SetState(432)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(435)
			p.Match(CQL3ParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.Column_name()
		}
		{
			p.SetState(438)
			p.Match(CQL3ParserK_IN)
		}
		{
			p.SetState(439)
			p.Match(CQL3ParserT__9)
		}

	}

	return localctx
}

// IDelete_stmtContext is an interface to support dynamic dispatch.
type IDelete_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_stmtContext differentiates from other interfaces.
	IsDelete_stmtContext()
}

type Delete_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_stmtContext() *Delete_stmtContext {
	var p = new(Delete_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_delete_stmt
	return p
}

func (*Delete_stmtContext) IsDelete_stmtContext() {}

func NewDelete_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_stmtContext {
	var p = new(Delete_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_delete_stmt

	return p
}

func (s *Delete_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_stmtContext) K_DELETE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_DELETE, 0)
}

func (s *Delete_stmtContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_FROM, 0)
}

func (s *Delete_stmtContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Delete_stmtContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_WHERE, 0)
}

func (s *Delete_stmtContext) Where_clause() IWhere_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhere_clauseContext)
}

func (s *Delete_stmtContext) Delete_selections() IDelete_selectionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_selectionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelete_selectionsContext)
}

func (s *Delete_stmtContext) K_USING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_USING, 0)
}

func (s *Delete_stmtContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TIMESTAMP, 0)
}

func (s *Delete_stmtContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserINTEGER, 0)
}

func (s *Delete_stmtContext) Delete_conditions() IDelete_conditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_conditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelete_conditionsContext)
}

func (s *Delete_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDelete_stmt(s)
	}
}

func (s *Delete_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDelete_stmt(s)
	}
}

func (p *CQL3Parser) Delete_stmt() (localctx IDelete_stmtContext) {
	localctx = NewDelete_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CQL3ParserRULE_delete_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(CQL3ParserK_DELETE)
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserIDENTIFIER {
		{
			p.SetState(444)
			p.Delete_selections()
		}

	}
	{
		p.SetState(447)
		p.Match(CQL3ParserK_FROM)
	}
	{
		p.SetState(448)
		p.Table_name()
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_USING {
		{
			p.SetState(449)
			p.Match(CQL3ParserK_USING)
		}
		{
			p.SetState(450)
			p.Match(CQL3ParserK_TIMESTAMP)
		}
		{
			p.SetState(451)
			p.Match(CQL3ParserINTEGER)
		}

	}
	{
		p.SetState(454)
		p.Match(CQL3ParserK_WHERE)
	}
	{
		p.SetState(455)
		p.Where_clause()
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_IF {
		{
			p.SetState(456)
			p.Delete_conditions()
		}

	}

	return localctx
}

// IDelete_conditionsContext is an interface to support dynamic dispatch.
type IDelete_conditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_conditionsContext differentiates from other interfaces.
	IsDelete_conditionsContext()
}

type Delete_conditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_conditionsContext() *Delete_conditionsContext {
	var p = new(Delete_conditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_delete_conditions
	return p
}

func (*Delete_conditionsContext) IsDelete_conditionsContext() {}

func NewDelete_conditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_conditionsContext {
	var p = new(Delete_conditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_delete_conditions

	return p
}

func (s *Delete_conditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_conditionsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_IF, 0)
}

func (s *Delete_conditionsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_EXISTS, 0)
}

func (s *Delete_conditionsContext) AllDelete_condition() []IDelete_conditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDelete_conditionContext)(nil)).Elem())
	var tst = make([]IDelete_conditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDelete_conditionContext)
		}
	}

	return tst
}

func (s *Delete_conditionsContext) Delete_condition(i int) IDelete_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_conditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDelete_conditionContext)
}

func (s *Delete_conditionsContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *Delete_conditionsContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *Delete_conditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_conditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_conditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDelete_conditions(s)
	}
}

func (s *Delete_conditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDelete_conditions(s)
	}
}

func (p *CQL3Parser) Delete_conditions() (localctx IDelete_conditionsContext) {
	localctx = NewDelete_conditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CQL3ParserRULE_delete_conditions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.Match(CQL3ParserK_IF)
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserK_EXISTS:
		{
			p.SetState(460)
			p.Match(CQL3ParserK_EXISTS)
		}

	case CQL3ParserIDENTIFIER:
		{
			p.SetState(461)
			p.Delete_condition()
		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQL3ParserK_AND {
			{
				p.SetState(462)
				p.Match(CQL3ParserK_AND)
			}
			{
				p.SetState(463)
				p.Delete_condition()
			}

			p.SetState(468)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDelete_conditionContext is an interface to support dynamic dispatch.
type IDelete_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_conditionContext differentiates from other interfaces.
	IsDelete_conditionContext()
}

type Delete_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_conditionContext() *Delete_conditionContext {
	var p = new(Delete_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_delete_condition
	return p
}

func (*Delete_conditionContext) IsDelete_conditionContext() {}

func NewDelete_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_conditionContext {
	var p = new(Delete_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_delete_condition

	return p
}

func (s *Delete_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_conditionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Delete_conditionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Delete_conditionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Delete_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDelete_condition(s)
	}
}

func (s *Delete_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDelete_condition(s)
	}
}

func (p *CQL3Parser) Delete_condition() (localctx IDelete_conditionContext) {
	localctx = NewDelete_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CQL3ParserRULE_delete_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Match(CQL3ParserIDENTIFIER)
	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserT__7 {
		{
			p.SetState(472)
			p.Match(CQL3ParserT__7)
		}
		{
			p.SetState(473)
			p.Term()
		}
		{
			p.SetState(474)
			p.Match(CQL3ParserT__8)
		}

	}
	{
		p.SetState(478)
		p.Match(CQL3ParserT__4)
	}
	{
		p.SetState(479)
		p.Term()
	}

	return localctx
}

// IDelete_selectionsContext is an interface to support dynamic dispatch.
type IDelete_selectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_selectionsContext differentiates from other interfaces.
	IsDelete_selectionsContext()
}

type Delete_selectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_selectionsContext() *Delete_selectionsContext {
	var p = new(Delete_selectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_delete_selections
	return p
}

func (*Delete_selectionsContext) IsDelete_selectionsContext() {}

func NewDelete_selectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_selectionsContext {
	var p = new(Delete_selectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_delete_selections

	return p
}

func (s *Delete_selectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_selectionsContext) AllDelete_selection() []IDelete_selectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDelete_selectionContext)(nil)).Elem())
	var tst = make([]IDelete_selectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDelete_selectionContext)
		}
	}

	return tst
}

func (s *Delete_selectionsContext) Delete_selection(i int) IDelete_selectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_selectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDelete_selectionContext)
}

func (s *Delete_selectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_selectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_selectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDelete_selections(s)
	}
}

func (s *Delete_selectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDelete_selections(s)
	}
}

func (p *CQL3Parser) Delete_selections() (localctx IDelete_selectionsContext) {
	localctx = NewDelete_selectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CQL3ParserRULE_delete_selections)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		p.Delete_selection()
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserT__3 {
		{
			p.SetState(482)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(483)
			p.Delete_selection()
		}

		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDelete_selectionContext is an interface to support dynamic dispatch.
type IDelete_selectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_selectionContext differentiates from other interfaces.
	IsDelete_selectionContext()
}

type Delete_selectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_selectionContext() *Delete_selectionContext {
	var p = new(Delete_selectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_delete_selection
	return p
}

func (*Delete_selectionContext) IsDelete_selectionContext() {}

func NewDelete_selectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_selectionContext {
	var p = new(Delete_selectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_delete_selection

	return p
}

func (s *Delete_selectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_selectionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Delete_selectionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Delete_selectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_selectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_selectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterDelete_selection(s)
	}
}

func (s *Delete_selectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitDelete_selection(s)
	}
}

func (p *CQL3Parser) Delete_selection() (localctx IDelete_selectionContext) {
	localctx = NewDelete_selectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CQL3ParserRULE_delete_selection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(CQL3ParserIDENTIFIER)
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserT__7 {
		{
			p.SetState(490)
			p.Match(CQL3ParserT__7)
		}
		{
			p.SetState(491)
			p.Term()
		}
		{
			p.SetState(492)
			p.Match(CQL3ParserT__8)
		}

	}

	return localctx
}

// IBatch_stmtContext is an interface to support dynamic dispatch.
type IBatch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatch_stmtContext differentiates from other interfaces.
	IsBatch_stmtContext()
}

type Batch_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatch_stmtContext() *Batch_stmtContext {
	var p = new(Batch_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_batch_stmt
	return p
}

func (*Batch_stmtContext) IsBatch_stmtContext() {}

func NewBatch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Batch_stmtContext {
	var p = new(Batch_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_batch_stmt

	return p
}

func (s *Batch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Batch_stmtContext) K_BEGIN() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_BEGIN, 0)
}

func (s *Batch_stmtContext) AllK_BATCH() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_BATCH)
}

func (s *Batch_stmtContext) K_BATCH(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_BATCH, i)
}

func (s *Batch_stmtContext) Dml_statements() IDml_statementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDml_statementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDml_statementsContext)
}

func (s *Batch_stmtContext) K_APPLY() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_APPLY, 0)
}

func (s *Batch_stmtContext) Batch_options() IBatch_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBatch_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBatch_optionsContext)
}

func (s *Batch_stmtContext) K_UNLOGGED() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_UNLOGGED, 0)
}

func (s *Batch_stmtContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_COUNTER, 0)
}

func (s *Batch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Batch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Batch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterBatch_stmt(s)
	}
}

func (s *Batch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitBatch_stmt(s)
	}
}

func (p *CQL3Parser) Batch_stmt() (localctx IBatch_stmtContext) {
	localctx = NewBatch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CQL3ParserRULE_batch_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Match(CQL3ParserK_BEGIN)
	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_COUNTER || _la == CQL3ParserK_UNLOGGED {
		p.SetState(497)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CQL3ParserK_COUNTER || _la == CQL3ParserK_UNLOGGED) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(500)
		p.Match(CQL3ParserK_BATCH)
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQL3ParserK_USING {
		{
			p.SetState(501)
			p.Batch_options()
		}

	}
	{
		p.SetState(504)
		p.Dml_statements()
	}
	{
		p.SetState(505)
		p.Match(CQL3ParserK_APPLY)
	}
	{
		p.SetState(506)
		p.Match(CQL3ParserK_BATCH)
	}

	return localctx
}

// IBatch_optionsContext is an interface to support dynamic dispatch.
type IBatch_optionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatch_optionsContext differentiates from other interfaces.
	IsBatch_optionsContext()
}

type Batch_optionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatch_optionsContext() *Batch_optionsContext {
	var p = new(Batch_optionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_batch_options
	return p
}

func (*Batch_optionsContext) IsBatch_optionsContext() {}

func NewBatch_optionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Batch_optionsContext {
	var p = new(Batch_optionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_batch_options

	return p
}

func (s *Batch_optionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Batch_optionsContext) K_USING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_USING, 0)
}

func (s *Batch_optionsContext) AllBatch_option() []IBatch_optionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBatch_optionContext)(nil)).Elem())
	var tst = make([]IBatch_optionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBatch_optionContext)
		}
	}

	return tst
}

func (s *Batch_optionsContext) Batch_option(i int) IBatch_optionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBatch_optionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBatch_optionContext)
}

func (s *Batch_optionsContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *Batch_optionsContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *Batch_optionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Batch_optionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Batch_optionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterBatch_options(s)
	}
}

func (s *Batch_optionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitBatch_options(s)
	}
}

func (p *CQL3Parser) Batch_options() (localctx IBatch_optionsContext) {
	localctx = NewBatch_optionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CQL3ParserRULE_batch_options)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Match(CQL3ParserK_USING)
	}
	{
		p.SetState(509)
		p.Batch_option()
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserK_AND {
		{
			p.SetState(510)
			p.Match(CQL3ParserK_AND)
		}
		{
			p.SetState(511)
			p.Batch_option()
		}

		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBatch_optionContext is an interface to support dynamic dispatch.
type IBatch_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatch_optionContext differentiates from other interfaces.
	IsBatch_optionContext()
}

type Batch_optionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatch_optionContext() *Batch_optionContext {
	var p = new(Batch_optionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_batch_option
	return p
}

func (*Batch_optionContext) IsBatch_optionContext() {}

func NewBatch_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Batch_optionContext {
	var p = new(Batch_optionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_batch_option

	return p
}

func (s *Batch_optionContext) GetParser() antlr.Parser { return s.parser }

func (s *Batch_optionContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TIMESTAMP, 0)
}

func (s *Batch_optionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserINTEGER, 0)
}

func (s *Batch_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Batch_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Batch_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterBatch_option(s)
	}
}

func (s *Batch_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitBatch_option(s)
	}
}

func (p *CQL3Parser) Batch_option() (localctx IBatch_optionContext) {
	localctx = NewBatch_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CQL3ParserRULE_batch_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(CQL3ParserK_TIMESTAMP)
	}
	{
		p.SetState(518)
		p.Match(CQL3ParserINTEGER)
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Table_nameContext) Keyspace_name() IKeyspace_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyspace_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyspace_nameContext)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterTable_name(s)
	}
}

func (s *Table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitTable_name(s)
	}
}

func (p *CQL3Parser) Table_name() (localctx ITable_nameContext) {
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CQL3ParserRULE_table_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(523)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(520)
			p.Keyspace_name()
		}
		{
			p.SetState(521)
			p.Match(CQL3ParserT__10)
		}

	}
	{
		p.SetState(525)
		p.Match(CQL3ParserIDENTIFIER)
	}

	return localctx
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_column_name
	return p
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterColumn_name(s)
	}
}

func (s *Column_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitColumn_name(s)
	}
}

func (p *CQL3Parser) Column_name() (localctx IColumn_nameContext) {
	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CQL3ParserRULE_column_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(527)
		p.Match(CQL3ParserIDENTIFIER)
	}

	return localctx
}

// ITable_optionsContext is an interface to support dynamic dispatch.
type ITable_optionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_optionsContext differentiates from other interfaces.
	IsTable_optionsContext()
}

type Table_optionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_optionsContext() *Table_optionsContext {
	var p = new(Table_optionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_table_options
	return p
}

func (*Table_optionsContext) IsTable_optionsContext() {}

func NewTable_optionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_optionsContext {
	var p = new(Table_optionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_table_options

	return p
}

func (s *Table_optionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_optionsContext) AllTable_option() []ITable_optionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITable_optionContext)(nil)).Elem())
	var tst = make([]ITable_optionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITable_optionContext)
		}
	}

	return tst
}

func (s *Table_optionsContext) Table_option(i int) ITable_optionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_optionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITable_optionContext)
}

func (s *Table_optionsContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *Table_optionsContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *Table_optionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_optionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_optionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterTable_options(s)
	}
}

func (s *Table_optionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitTable_options(s)
	}
}

func (p *CQL3Parser) Table_options() (localctx ITable_optionsContext) {
	localctx = NewTable_optionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CQL3ParserRULE_table_options)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		p.Table_option()
	}
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserK_AND {
		{
			p.SetState(530)
			p.Match(CQL3ParserK_AND)
		}
		{
			p.SetState(531)
			p.Table_option()
		}

		p.SetState(536)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITable_optionContext is an interface to support dynamic dispatch.
type ITable_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_optionContext differentiates from other interfaces.
	IsTable_optionContext()
}

type Table_optionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_optionContext() *Table_optionContext {
	var p = new(Table_optionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_table_option
	return p
}

func (*Table_optionContext) IsTable_optionContext() {}

func NewTable_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_optionContext {
	var p = new(Table_optionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_table_option

	return p
}

func (s *Table_optionContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_optionContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Table_optionContext) K_COMPACT() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_COMPACT, 0)
}

func (s *Table_optionContext) K_STORAGE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_STORAGE, 0)
}

func (s *Table_optionContext) K_CLUSTERING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_CLUSTERING, 0)
}

func (s *Table_optionContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_ORDER, 0)
}

func (s *Table_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterTable_option(s)
	}
}

func (s *Table_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitTable_option(s)
	}
}

func (p *CQL3Parser) Table_option() (localctx ITable_optionContext) {
	localctx = NewTable_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CQL3ParserRULE_table_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(542)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			p.Property()
		}

	case CQL3ParserK_COMPACT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(538)
			p.Match(CQL3ParserK_COMPACT)
		}
		{
			p.SetState(539)
			p.Match(CQL3ParserK_STORAGE)
		}

	case CQL3ParserK_CLUSTERING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(540)
			p.Match(CQL3ParserK_CLUSTERING)
		}
		{
			p.SetState(541)
			p.Match(CQL3ParserK_ORDER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumn_definitionsContext is an interface to support dynamic dispatch.
type IColumn_definitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_definitionsContext differentiates from other interfaces.
	IsColumn_definitionsContext()
}

type Column_definitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_definitionsContext() *Column_definitionsContext {
	var p = new(Column_definitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_column_definitions
	return p
}

func (*Column_definitionsContext) IsColumn_definitionsContext() {}

func NewColumn_definitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_definitionsContext {
	var p = new(Column_definitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_column_definitions

	return p
}

func (s *Column_definitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_definitionsContext) AllColumn_definition() []IColumn_definitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_definitionContext)(nil)).Elem())
	var tst = make([]IColumn_definitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_definitionContext)
		}
	}

	return tst
}

func (s *Column_definitionsContext) Column_definition(i int) IColumn_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_definitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_definitionContext)
}

func (s *Column_definitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_definitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_definitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterColumn_definitions(s)
	}
}

func (s *Column_definitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitColumn_definitions(s)
	}
}

func (p *CQL3Parser) Column_definitions() (localctx IColumn_definitionsContext) {
	localctx = NewColumn_definitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CQL3ParserRULE_column_definitions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(544)
		p.Match(CQL3ParserT__1)
	}
	{
		p.SetState(545)
		p.Column_definition()
	}
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserT__3 {
		{
			p.SetState(546)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(547)
			p.Column_definition()
		}

		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(553)
		p.Match(CQL3ParserT__2)
	}

	return localctx
}

// IColumn_definitionContext is an interface to support dynamic dispatch.
type IColumn_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_definitionContext differentiates from other interfaces.
	IsColumn_definitionContext()
}

type Column_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_definitionContext() *Column_definitionContext {
	var p = new(Column_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_column_definition
	return p
}

func (*Column_definitionContext) IsColumn_definitionContext() {}

func NewColumn_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_definitionContext {
	var p = new(Column_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_column_definition

	return p
}

func (s *Column_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_definitionContext) Column_name() IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_definitionContext) Column_type() IColumn_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_typeContext)
}

func (s *Column_definitionContext) K_STATIC() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_STATIC, 0)
}

func (s *Column_definitionContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_PRIMARY, 0)
}

func (s *Column_definitionContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_KEY, 0)
}

func (s *Column_definitionContext) Primary_key() IPrimary_keyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_keyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_keyContext)
}

func (s *Column_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterColumn_definition(s)
	}
}

func (s *Column_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitColumn_definition(s)
	}
}

func (p *CQL3Parser) Column_definition() (localctx IColumn_definitionContext) {
	localctx = NewColumn_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, CQL3ParserRULE_column_definition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(567)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)
			p.Column_name()
		}
		{
			p.SetState(556)
			p.Column_type()
		}
		p.SetState(558)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CQL3ParserK_STATIC {
			{
				p.SetState(557)
				p.Match(CQL3ParserK_STATIC)
			}

		}
		p.SetState(562)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CQL3ParserK_PRIMARY {
			{
				p.SetState(560)
				p.Match(CQL3ParserK_PRIMARY)
			}
			{
				p.SetState(561)
				p.Match(CQL3ParserK_KEY)
			}

		}

	case CQL3ParserK_PRIMARY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(564)
			p.Match(CQL3ParserK_PRIMARY)
		}
		{
			p.SetState(565)
			p.Match(CQL3ParserK_KEY)
		}
		{
			p.SetState(566)
			p.Primary_key()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumn_typeContext is an interface to support dynamic dispatch.
type IColumn_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_typeContext differentiates from other interfaces.
	IsColumn_typeContext()
}

type Column_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_typeContext() *Column_typeContext {
	var p = new(Column_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_column_type
	return p
}

func (*Column_typeContext) IsColumn_typeContext() {}

func NewColumn_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_typeContext {
	var p = new(Column_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_column_type

	return p
}

func (s *Column_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_typeContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Column_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterColumn_type(s)
	}
}

func (s *Column_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitColumn_type(s)
	}
}

func (p *CQL3Parser) Column_type() (localctx IColumn_typeContext) {
	localctx = NewColumn_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, CQL3ParserRULE_column_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Data_type()
	}

	return localctx
}

// IPrimary_keyContext is an interface to support dynamic dispatch.
type IPrimary_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimary_keyContext differentiates from other interfaces.
	IsPrimary_keyContext()
}

type Primary_keyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_keyContext() *Primary_keyContext {
	var p = new(Primary_keyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_primary_key
	return p
}

func (*Primary_keyContext) IsPrimary_keyContext() {}

func NewPrimary_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_keyContext {
	var p = new(Primary_keyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_primary_key

	return p
}

func (s *Primary_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_keyContext) Partition_key() IPartition_keyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartition_keyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartition_keyContext)
}

func (s *Primary_keyContext) AllColumn_name() []IColumn_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem())
	var tst = make([]IColumn_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_nameContext)
		}
	}

	return tst
}

func (s *Primary_keyContext) Column_name(i int) IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Primary_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterPrimary_key(s)
	}
}

func (s *Primary_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitPrimary_key(s)
	}
}

func (p *CQL3Parser) Primary_key() (localctx IPrimary_keyContext) {
	localctx = NewPrimary_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, CQL3ParserRULE_primary_key)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)
		p.Match(CQL3ParserT__1)
	}
	{
		p.SetState(572)
		p.Partition_key()
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserT__3 {
		{
			p.SetState(573)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(574)
			p.Column_name()
		}

		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(580)
		p.Match(CQL3ParserT__2)
	}

	return localctx
}

// IPartition_keyContext is an interface to support dynamic dispatch.
type IPartition_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartition_keyContext differentiates from other interfaces.
	IsPartition_keyContext()
}

type Partition_keyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartition_keyContext() *Partition_keyContext {
	var p = new(Partition_keyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_partition_key
	return p
}

func (*Partition_keyContext) IsPartition_keyContext() {}

func NewPartition_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Partition_keyContext {
	var p = new(Partition_keyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_partition_key

	return p
}

func (s *Partition_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Partition_keyContext) AllColumn_name() []IColumn_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem())
	var tst = make([]IColumn_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_nameContext)
		}
	}

	return tst
}

func (s *Partition_keyContext) Column_name(i int) IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Partition_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Partition_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Partition_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterPartition_key(s)
	}
}

func (s *Partition_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitPartition_key(s)
	}
}

func (p *CQL3Parser) Partition_key() (localctx IPartition_keyContext) {
	localctx = NewPartition_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, CQL3ParserRULE_partition_key)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(594)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(582)
			p.Column_name()
		}

	case CQL3ParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(583)
			p.Match(CQL3ParserT__1)
		}
		{
			p.SetState(584)
			p.Column_name()
		}
		p.SetState(589)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQL3ParserT__3 {
			{
				p.SetState(585)
				p.Match(CQL3ParserT__3)
			}
			{
				p.SetState(586)
				p.Column_name()
			}

			p.SetState(591)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(592)
			p.Match(CQL3ParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeyspace_nameContext is an interface to support dynamic dispatch.
type IKeyspace_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyspace_nameContext differentiates from other interfaces.
	IsKeyspace_nameContext()
}

type Keyspace_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyspace_nameContext() *Keyspace_nameContext {
	var p = new(Keyspace_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_keyspace_name
	return p
}

func (*Keyspace_nameContext) IsKeyspace_nameContext() {}

func NewKeyspace_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Keyspace_nameContext {
	var p = new(Keyspace_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_keyspace_name

	return p
}

func (s *Keyspace_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Keyspace_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Keyspace_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Keyspace_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Keyspace_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterKeyspace_name(s)
	}
}

func (s *Keyspace_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitKeyspace_name(s)
	}
}

func (p *CQL3Parser) Keyspace_name() (localctx IKeyspace_nameContext) {
	localctx = NewKeyspace_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, CQL3ParserRULE_keyspace_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		p.Match(CQL3ParserIDENTIFIER)
	}

	return localctx
}

// IIf_not_existsContext is an interface to support dynamic dispatch.
type IIf_not_existsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_not_existsContext differentiates from other interfaces.
	IsIf_not_existsContext()
}

type If_not_existsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_not_existsContext() *If_not_existsContext {
	var p = new(If_not_existsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_if_not_exists
	return p
}

func (*If_not_existsContext) IsIf_not_existsContext() {}

func NewIf_not_existsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_not_existsContext {
	var p = new(If_not_existsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_if_not_exists

	return p
}

func (s *If_not_existsContext) GetParser() antlr.Parser { return s.parser }

func (s *If_not_existsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_IF, 0)
}

func (s *If_not_existsContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_NOT, 0)
}

func (s *If_not_existsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_EXISTS, 0)
}

func (s *If_not_existsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_not_existsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_not_existsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterIf_not_exists(s)
	}
}

func (s *If_not_existsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitIf_not_exists(s)
	}
}

func (p *CQL3Parser) If_not_exists() (localctx IIf_not_existsContext) {
	localctx = NewIf_not_existsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, CQL3ParserRULE_if_not_exists)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(598)
		p.Match(CQL3ParserK_IF)
	}
	{
		p.SetState(599)
		p.Match(CQL3ParserK_NOT)
	}
	{
		p.SetState(600)
		p.Match(CQL3ParserK_EXISTS)
	}

	return localctx
}

// IIf_existsContext is an interface to support dynamic dispatch.
type IIf_existsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_existsContext differentiates from other interfaces.
	IsIf_existsContext()
}

type If_existsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_existsContext() *If_existsContext {
	var p = new(If_existsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_if_exists
	return p
}

func (*If_existsContext) IsIf_existsContext() {}

func NewIf_existsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_existsContext {
	var p = new(If_existsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_if_exists

	return p
}

func (s *If_existsContext) GetParser() antlr.Parser { return s.parser }

func (s *If_existsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_IF, 0)
}

func (s *If_existsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_EXISTS, 0)
}

func (s *If_existsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_existsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_existsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterIf_exists(s)
	}
}

func (s *If_existsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitIf_exists(s)
	}
}

func (p *CQL3Parser) If_exists() (localctx IIf_existsContext) {
	localctx = NewIf_existsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, CQL3ParserRULE_if_exists)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(602)
		p.Match(CQL3ParserK_IF)
	}
	{
		p.SetState(603)
		p.Match(CQL3ParserK_EXISTS)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) STRING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserSTRING, 0)
}

func (s *ConstantContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserINTEGER, 0)
}

func (s *ConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CQL3ParserFLOAT, 0)
}

func (s *ConstantContext) Bool() IBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *ConstantContext) UUID() antlr.TerminalNode {
	return s.GetToken(CQL3ParserUUID, 0)
}

func (s *ConstantContext) BLOB() antlr.TerminalNode {
	return s.GetToken(CQL3ParserBLOB, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *CQL3Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, CQL3ParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(611)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(605)
			p.Match(CQL3ParserSTRING)
		}

	case CQL3ParserINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(606)
			p.Match(CQL3ParserINTEGER)
		}

	case CQL3ParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(607)
			p.Match(CQL3ParserFLOAT)
		}

	case CQL3ParserK_FALSE, CQL3ParserK_TRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(608)
			p.Bool()
		}

	case CQL3ParserUUID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(609)
			p.Match(CQL3ParserUUID)
		}

	case CQL3ParserBLOB:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(610)
			p.Match(CQL3ParserBLOB)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *CQL3Parser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, CQL3ParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(616)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(613)
			p.Match(CQL3ParserT__9)
		}

	case CQL3ParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(614)
			p.Match(CQL3ParserT__11)
		}
		{
			p.SetState(615)
			p.Match(CQL3ParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *TermContext) Collection() ICollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *TermContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *TermContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *CQL3Parser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, CQL3ParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(622)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserK_FALSE, CQL3ParserK_TRUE, CQL3ParserSTRING, CQL3ParserINTEGER, CQL3ParserFLOAT, CQL3ParserUUID, CQL3ParserBLOB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(618)
			p.Constant()
		}

	case CQL3ParserT__7, CQL3ParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(619)
			p.Collection()
		}

	case CQL3ParserT__9, CQL3ParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(620)
			p.Variable()
		}

	case CQL3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(621)
			p.Function()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_collection
	return p
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) Map() IMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *CollectionContext) Set() ISetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *CollectionContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitCollection(s)
	}
}

func (p *CQL3Parser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, CQL3ParserRULE_collection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(624)
			p.Map()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(625)
			p.Set()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(626)
			p.List()
		}

	}

	return localctx
}

// IMapContext is an interface to support dynamic dispatch.
type IMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapContext differentiates from other interfaces.
	IsMapContext()
}

type MapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapContext() *MapContext {
	var p = new(MapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_map
	return p
}

func (*MapContext) IsMapContext() {}

func NewMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapContext {
	var p = new(MapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_map

	return p
}

func (s *MapContext) GetParser() antlr.Parser { return s.parser }

func (s *MapContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *MapContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitMap(s)
	}
}

func (p *CQL3Parser) Map() (localctx IMapContext) {
	localctx = NewMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, CQL3ParserRULE_map)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(629)
		p.Match(CQL3ParserT__12)
	}
	p.SetState(643)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQL3ParserT__7)|(1<<CQL3ParserT__9)|(1<<CQL3ParserT__11)|(1<<CQL3ParserT__12))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(CQL3ParserK_FALSE-51))|(1<<(CQL3ParserK_TRUE-51))|(1<<(CQL3ParserIDENTIFIER-51)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(CQL3ParserSTRING-83))|(1<<(CQL3ParserINTEGER-83))|(1<<(CQL3ParserFLOAT-83))|(1<<(CQL3ParserUUID-83))|(1<<(CQL3ParserBLOB-83)))) != 0) {
		{
			p.SetState(630)
			p.Term()
		}
		{
			p.SetState(631)
			p.Match(CQL3ParserT__11)
		}
		{
			p.SetState(632)
			p.Term()
		}
		p.SetState(640)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQL3ParserT__3 {
			{
				p.SetState(633)
				p.Match(CQL3ParserT__3)
			}
			{
				p.SetState(634)
				p.Term()
			}
			{
				p.SetState(635)
				p.Match(CQL3ParserT__11)
			}
			{
				p.SetState(636)
				p.Term()
			}

			p.SetState(642)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(645)
		p.Match(CQL3ParserT__13)
	}

	return localctx
}

// ISetContext is an interface to support dynamic dispatch.
type ISetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetContext differentiates from other interfaces.
	IsSetContext()
}

type SetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetContext() *SetContext {
	var p = new(SetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_set
	return p
}

func (*SetContext) IsSetContext() {}

func NewSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetContext {
	var p = new(SetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_set

	return p
}

func (s *SetContext) GetParser() antlr.Parser { return s.parser }

func (s *SetContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *SetContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterSet(s)
	}
}

func (s *SetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitSet(s)
	}
}

func (p *CQL3Parser) Set() (localctx ISetContext) {
	localctx = NewSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, CQL3ParserRULE_set)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)
		p.Match(CQL3ParserT__12)
	}
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQL3ParserT__7)|(1<<CQL3ParserT__9)|(1<<CQL3ParserT__11)|(1<<CQL3ParserT__12))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(CQL3ParserK_FALSE-51))|(1<<(CQL3ParserK_TRUE-51))|(1<<(CQL3ParserIDENTIFIER-51)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(CQL3ParserSTRING-83))|(1<<(CQL3ParserINTEGER-83))|(1<<(CQL3ParserFLOAT-83))|(1<<(CQL3ParserUUID-83))|(1<<(CQL3ParserBLOB-83)))) != 0) {
		{
			p.SetState(648)
			p.Term()
		}
		p.SetState(653)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQL3ParserT__3 {
			{
				p.SetState(649)
				p.Match(CQL3ParserT__3)
			}
			{
				p.SetState(650)
				p.Term()
			}

			p.SetState(655)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(658)
		p.Match(CQL3ParserT__13)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ListContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitList(s)
	}
}

func (p *CQL3Parser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, CQL3ParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(660)
		p.Match(CQL3ParserT__7)
	}
	p.SetState(669)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQL3ParserT__7)|(1<<CQL3ParserT__9)|(1<<CQL3ParserT__11)|(1<<CQL3ParserT__12))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(CQL3ParserK_FALSE-51))|(1<<(CQL3ParserK_TRUE-51))|(1<<(CQL3ParserIDENTIFIER-51)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(CQL3ParserSTRING-83))|(1<<(CQL3ParserINTEGER-83))|(1<<(CQL3ParserFLOAT-83))|(1<<(CQL3ParserUUID-83))|(1<<(CQL3ParserBLOB-83)))) != 0) {
		{
			p.SetState(661)
			p.Term()
		}
		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQL3ParserT__3 {
			{
				p.SetState(662)
				p.Match(CQL3ParserT__3)
			}
			{
				p.SetState(663)
				p.Term()
			}

			p.SetState(668)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(671)
		p.Match(CQL3ParserT__8)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *FunctionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FunctionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *CQL3Parser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, CQL3ParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(673)
		p.Match(CQL3ParserIDENTIFIER)
	}
	{
		p.SetState(674)
		p.Match(CQL3ParserT__1)
	}
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQL3ParserT__7)|(1<<CQL3ParserT__9)|(1<<CQL3ParserT__11)|(1<<CQL3ParserT__12))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(CQL3ParserK_FALSE-51))|(1<<(CQL3ParserK_TRUE-51))|(1<<(CQL3ParserIDENTIFIER-51)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(CQL3ParserSTRING-83))|(1<<(CQL3ParserINTEGER-83))|(1<<(CQL3ParserFLOAT-83))|(1<<(CQL3ParserUUID-83))|(1<<(CQL3ParserBLOB-83)))) != 0) {
		{
			p.SetState(675)
			p.Term()
		}
		p.SetState(680)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQL3ParserT__3 {
			{
				p.SetState(676)
				p.Match(CQL3ParserT__3)
			}
			{
				p.SetState(677)
				p.Term()
			}

			p.SetState(682)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(685)
		p.Match(CQL3ParserT__2)
	}

	return localctx
}

// IPropertiesContext is an interface to support dynamic dispatch.
type IPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertiesContext differentiates from other interfaces.
	IsPropertiesContext()
}

type PropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesContext() *PropertiesContext {
	var p = new(PropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_properties
	return p
}

func (*PropertiesContext) IsPropertiesContext() {}

func NewPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesContext {
	var p = new(PropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_properties

	return p
}

func (s *PropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *PropertiesContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertiesContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(CQL3ParserK_AND)
}

func (s *PropertiesContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_AND, i)
}

func (s *PropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterProperties(s)
	}
}

func (s *PropertiesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitProperties(s)
	}
}

func (p *CQL3Parser) Properties() (localctx IPropertiesContext) {
	localctx = NewPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, CQL3ParserRULE_properties)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(687)
		p.Property()
	}
	p.SetState(692)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQL3ParserK_AND {
		{
			p.SetState(688)
			p.Match(CQL3ParserK_AND)
		}
		{
			p.SetState(689)
			p.Property()
		}

		p.SetState(694)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Property_name() IProperty_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *PropertyContext) Property_value() IProperty_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProperty_valueContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *CQL3Parser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, CQL3ParserRULE_property)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(695)
		p.Property_name()
	}
	{
		p.SetState(696)
		p.Match(CQL3ParserT__4)
	}
	{
		p.SetState(697)
		p.Property_value()
	}

	return localctx
}

// IProperty_nameContext is an interface to support dynamic dispatch.
type IProperty_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProperty_nameContext differentiates from other interfaces.
	IsProperty_nameContext()
}

type Property_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_nameContext() *Property_nameContext {
	var p = new(Property_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_property_name
	return p
}

func (*Property_nameContext) IsProperty_nameContext() {}

func NewProperty_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_nameContext {
	var p = new(Property_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_property_name

	return p
}

func (s *Property_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Property_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterProperty_name(s)
	}
}

func (s *Property_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitProperty_name(s)
	}
}

func (p *CQL3Parser) Property_name() (localctx IProperty_nameContext) {
	localctx = NewProperty_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, CQL3ParserRULE_property_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(699)
		p.Match(CQL3ParserIDENTIFIER)
	}

	return localctx
}

// IProperty_valueContext is an interface to support dynamic dispatch.
type IProperty_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProperty_valueContext differentiates from other interfaces.
	IsProperty_valueContext()
}

type Property_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_valueContext() *Property_valueContext {
	var p = new(Property_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_property_value
	return p
}

func (*Property_valueContext) IsProperty_valueContext() {}

func NewProperty_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_valueContext {
	var p = new(Property_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_property_value

	return p
}

func (s *Property_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_valueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CQL3ParserIDENTIFIER, 0)
}

func (s *Property_valueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Property_valueContext) Map() IMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *Property_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterProperty_value(s)
	}
}

func (s *Property_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitProperty_value(s)
	}
}

func (p *CQL3Parser) Property_value() (localctx IProperty_valueContext) {
	localctx = NewProperty_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, CQL3ParserRULE_property_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(704)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(701)
			p.Match(CQL3ParserIDENTIFIER)
		}

	case CQL3ParserK_FALSE, CQL3ParserK_TRUE, CQL3ParserSTRING, CQL3ParserINTEGER, CQL3ParserFLOAT, CQL3ParserUUID, CQL3ParserBLOB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(702)
			p.Constant()
		}

	case CQL3ParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(703)
			p.Map()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_data_type
	return p
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) Native_type() INative_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INative_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INative_typeContext)
}

func (s *Data_typeContext) Collection_type() ICollection_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_typeContext)
}

func (s *Data_typeContext) STRING() antlr.TerminalNode {
	return s.GetToken(CQL3ParserSTRING, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterData_type(s)
	}
}

func (s *Data_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitData_type(s)
	}
}

func (p *CQL3Parser) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, CQL3ParserRULE_data_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(709)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserT__14, CQL3ParserT__15, CQL3ParserT__16, CQL3ParserT__17, CQL3ParserT__18, CQL3ParserT__19, CQL3ParserT__20, CQL3ParserT__21, CQL3ParserT__22, CQL3ParserT__23, CQL3ParserT__24, CQL3ParserT__25, CQL3ParserT__26, CQL3ParserT__27, CQL3ParserT__28, CQL3ParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(706)
			p.Native_type()
		}

	case CQL3ParserT__30, CQL3ParserT__33, CQL3ParserT__34:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(707)
			p.Collection_type()
		}

	case CQL3ParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(708)
			p.Match(CQL3ParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INative_typeContext is an interface to support dynamic dispatch.
type INative_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNative_typeContext differentiates from other interfaces.
	IsNative_typeContext()
}

type Native_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNative_typeContext() *Native_typeContext {
	var p = new(Native_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_native_type
	return p
}

func (*Native_typeContext) IsNative_typeContext() {}

func NewNative_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Native_typeContext {
	var p = new(Native_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_native_type

	return p
}

func (s *Native_typeContext) GetParser() antlr.Parser { return s.parser }
func (s *Native_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Native_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Native_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterNative_type(s)
	}
}

func (s *Native_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitNative_type(s)
	}
}

func (p *CQL3Parser) Native_type() (localctx INative_typeContext) {
	localctx = NewNative_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, CQL3ParserRULE_native_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(711)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQL3ParserT__14)|(1<<CQL3ParserT__15)|(1<<CQL3ParserT__16)|(1<<CQL3ParserT__17)|(1<<CQL3ParserT__18)|(1<<CQL3ParserT__19)|(1<<CQL3ParserT__20)|(1<<CQL3ParserT__21)|(1<<CQL3ParserT__22)|(1<<CQL3ParserT__23)|(1<<CQL3ParserT__24)|(1<<CQL3ParserT__25)|(1<<CQL3ParserT__26)|(1<<CQL3ParserT__27)|(1<<CQL3ParserT__28)|(1<<CQL3ParserT__29))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ICollection_typeContext is an interface to support dynamic dispatch.
type ICollection_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollection_typeContext differentiates from other interfaces.
	IsCollection_typeContext()
}

type Collection_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_typeContext() *Collection_typeContext {
	var p = new(Collection_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_collection_type
	return p
}

func (*Collection_typeContext) IsCollection_typeContext() {}

func NewCollection_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_typeContext {
	var p = new(Collection_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_collection_type

	return p
}

func (s *Collection_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_typeContext) AllNative_type() []INative_typeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INative_typeContext)(nil)).Elem())
	var tst = make([]INative_typeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INative_typeContext)
		}
	}

	return tst
}

func (s *Collection_typeContext) Native_type(i int) INative_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INative_typeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INative_typeContext)
}

func (s *Collection_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collection_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterCollection_type(s)
	}
}

func (s *Collection_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitCollection_type(s)
	}
}

func (p *CQL3Parser) Collection_type() (localctx ICollection_typeContext) {
	localctx = NewCollection_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, CQL3ParserRULE_collection_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(730)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQL3ParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(713)
			p.Match(CQL3ParserT__30)
		}
		{
			p.SetState(714)
			p.Match(CQL3ParserT__31)
		}
		{
			p.SetState(715)
			p.Native_type()
		}
		{
			p.SetState(716)
			p.Match(CQL3ParserT__32)
		}

	case CQL3ParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(718)
			p.Match(CQL3ParserT__33)
		}
		{
			p.SetState(719)
			p.Match(CQL3ParserT__31)
		}
		{
			p.SetState(720)
			p.Native_type()
		}
		{
			p.SetState(721)
			p.Match(CQL3ParserT__32)
		}

	case CQL3ParserT__34:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(723)
			p.Match(CQL3ParserT__34)
		}
		{
			p.SetState(724)
			p.Match(CQL3ParserT__31)
		}
		{
			p.SetState(725)
			p.Native_type()
		}
		{
			p.SetState(726)
			p.Match(CQL3ParserT__3)
		}
		{
			p.SetState(727)
			p.Native_type()
		}
		{
			p.SetState(728)
			p.Match(CQL3ParserT__32)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQL3ParserRULE_bool
	return p
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQL3ParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) K_TRUE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_TRUE, 0)
}

func (s *BoolContext) K_FALSE() antlr.TerminalNode {
	return s.GetToken(CQL3ParserK_FALSE, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQL3Listener); ok {
		listenerT.ExitBool(s)
	}
}

func (p *CQL3Parser) Bool() (localctx IBoolContext) {
	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, CQL3ParserRULE_bool)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(732)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQL3ParserK_FALSE || _la == CQL3ParserK_TRUE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
