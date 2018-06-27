// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 690,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 5, 2, 75, 10, 2, 3, 2, 7,
	2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 6, 3, 87,
	10, 3, 13, 3, 14, 3, 88, 3, 3, 3, 3, 6, 3, 93, 10, 3, 13, 3, 14, 3, 94,
	3, 4, 3, 4, 3, 4, 6, 4, 100, 10, 4, 13, 4, 14, 4, 101, 3, 4, 7, 4, 105,
	10, 4, 12, 4, 14, 4, 108, 11, 4, 3, 5, 7, 5, 111, 10, 5, 12, 5, 14, 5,
	114, 11, 5, 3, 5, 3, 5, 3, 5, 7, 5, 119, 10, 5, 12, 5, 14, 5, 122, 11,
	5, 3, 5, 3, 5, 7, 5, 126, 10, 5, 12, 5, 14, 5, 129, 11, 5, 3, 5, 3, 5,
	7, 5, 133, 10, 5, 12, 5, 14, 5, 136, 11, 5, 3, 5, 3, 5, 5, 5, 140, 10,
	5, 3, 6, 7, 6, 143, 10, 6, 12, 6, 14, 6, 146, 11, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 5, 7, 153, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 158, 10, 8, 12, 8,
	14, 8, 161, 11, 8, 3, 8, 5, 8, 164, 10, 8, 3, 8, 3, 8, 7, 8, 168, 10, 8,
	12, 8, 14, 8, 171, 11, 8, 3, 9, 3, 9, 7, 9, 175, 10, 9, 12, 9, 14, 9, 178,
	11, 9, 3, 9, 3, 9, 3, 9, 7, 9, 183, 10, 9, 12, 9, 14, 9, 186, 11, 9, 3,
	9, 5, 9, 189, 10, 9, 3, 9, 3, 9, 7, 9, 193, 10, 9, 12, 9, 14, 9, 196, 11,
	9, 3, 10, 7, 10, 199, 10, 10, 12, 10, 14, 10, 202, 11, 10, 3, 10, 3, 10,
	7, 10, 206, 10, 10, 12, 10, 14, 10, 209, 11, 10, 3, 10, 3, 10, 3, 10, 7,
	10, 214, 10, 10, 12, 10, 14, 10, 217, 11, 10, 7, 10, 219, 10, 10, 12, 10,
	14, 10, 222, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 229, 10,
	11, 3, 12, 3, 12, 6, 12, 233, 10, 12, 13, 12, 14, 12, 234, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 246, 10, 14, 12,
	14, 14, 14, 249, 11, 14, 3, 14, 3, 14, 7, 14, 253, 10, 14, 12, 14, 14,
	14, 256, 11, 14, 3, 14, 7, 14, 259, 10, 14, 12, 14, 14, 14, 262, 11, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 268, 10, 15, 12, 15, 14, 15, 271, 11,
	15, 3, 15, 3, 15, 7, 15, 275, 10, 15, 12, 15, 14, 15, 278, 11, 15, 3, 15,
	3, 15, 7, 15, 282, 10, 15, 12, 15, 14, 15, 285, 11, 15, 3, 15, 3, 15, 7,
	15, 289, 10, 15, 12, 15, 14, 15, 292, 11, 15, 7, 15, 294, 10, 15, 12, 15,
	14, 15, 297, 11, 15, 5, 15, 299, 10, 15, 3, 15, 7, 15, 302, 10, 15, 12,
	15, 14, 15, 305, 11, 15, 3, 15, 5, 15, 308, 10, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 320, 10, 15, 12,
	15, 14, 15, 323, 11, 15, 3, 16, 3, 16, 7, 16, 327, 10, 16, 12, 16, 14,
	16, 330, 11, 16, 3, 16, 3, 16, 5, 16, 334, 10, 16, 3, 16, 7, 16, 337, 10,
	16, 12, 16, 14, 16, 340, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 7, 17, 348, 10, 17, 12, 17, 14, 17, 351, 11, 17, 3, 17, 3, 17, 7, 17,
	355, 10, 17, 12, 17, 14, 17, 358, 11, 17, 3, 17, 3, 17, 3, 17, 5, 17, 363,
	10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 371, 10, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 6, 22, 379, 10, 22, 13, 22, 14,
	22, 380, 3, 22, 3, 22, 7, 22, 385, 10, 22, 12, 22, 14, 22, 388, 11, 22,
	3, 22, 3, 22, 5, 22, 392, 10, 22, 3, 22, 3, 22, 6, 22, 396, 10, 22, 13,
	22, 14, 22, 397, 3, 22, 3, 22, 7, 22, 402, 10, 22, 12, 22, 14, 22, 405,
	11, 22, 3, 22, 3, 22, 7, 22, 409, 10, 22, 12, 22, 14, 22, 412, 11, 22,
	3, 22, 5, 22, 415, 10, 22, 3, 23, 3, 23, 6, 23, 419, 10, 23, 13, 23, 14,
	23, 420, 3, 24, 3, 24, 5, 24, 425, 10, 24, 3, 25, 3, 25, 6, 25, 429, 10,
	25, 13, 25, 14, 25, 430, 3, 25, 3, 25, 7, 25, 435, 10, 25, 12, 25, 14,
	25, 438, 11, 25, 3, 25, 3, 25, 7, 25, 442, 10, 25, 12, 25, 14, 25, 445,
	11, 25, 3, 25, 3, 25, 7, 25, 449, 10, 25, 12, 25, 14, 25, 452, 11, 25,
	6, 25, 454, 10, 25, 13, 25, 14, 25, 455, 5, 25, 458, 10, 25, 3, 25, 7,
	25, 461, 10, 25, 12, 25, 14, 25, 464, 11, 25, 3, 25, 5, 25, 467, 10, 25,
	3, 25, 7, 25, 470, 10, 25, 12, 25, 14, 25, 473, 11, 25, 3, 25, 3, 25, 7,
	25, 477, 10, 25, 12, 25, 14, 25, 480, 11, 25, 3, 25, 7, 25, 483, 10, 25,
	12, 25, 14, 25, 486, 11, 25, 3, 25, 7, 25, 489, 10, 25, 12, 25, 14, 25,
	492, 11, 25, 3, 25, 3, 25, 3, 26, 7, 26, 497, 10, 26, 12, 26, 14, 26, 500,
	11, 26, 3, 26, 3, 26, 7, 26, 504, 10, 26, 12, 26, 14, 26, 507, 11, 26,
	3, 26, 3, 26, 6, 26, 511, 10, 26, 13, 26, 14, 26, 512, 3, 27, 3, 27, 6,
	27, 517, 10, 27, 13, 27, 14, 27, 518, 3, 27, 3, 27, 7, 27, 523, 10, 27,
	12, 27, 14, 27, 526, 11, 27, 3, 27, 5, 27, 529, 10, 27, 3, 27, 7, 27, 532,
	10, 27, 12, 27, 14, 27, 535, 11, 27, 3, 27, 3, 27, 7, 27, 539, 10, 27,
	12, 27, 14, 27, 542, 11, 27, 3, 27, 3, 27, 6, 27, 546, 10, 27, 13, 27,
	14, 27, 547, 7, 27, 550, 10, 27, 12, 27, 14, 27, 553, 11, 27, 3, 27, 7,
	27, 556, 10, 27, 12, 27, 14, 27, 559, 11, 27, 3, 27, 3, 27, 3, 28, 7, 28,
	564, 10, 28, 12, 28, 14, 28, 567, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3,
	29, 7, 29, 574, 10, 29, 12, 29, 14, 29, 577, 11, 29, 3, 29, 3, 29, 3, 29,
	7, 29, 582, 10, 29, 12, 29, 14, 29, 585, 11, 29, 5, 29, 587, 10, 29, 3,
	29, 7, 29, 590, 10, 29, 12, 29, 14, 29, 593, 11, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 7, 31, 602, 10, 31, 12, 31, 14, 31, 605, 11,
	31, 3, 31, 3, 31, 3, 31, 7, 31, 610, 10, 31, 12, 31, 14, 31, 613, 11, 31,
	5, 31, 615, 10, 31, 3, 31, 3, 31, 3, 32, 7, 32, 620, 10, 32, 12, 32, 14,
	32, 623, 11, 32, 3, 32, 3, 32, 7, 32, 627, 10, 32, 12, 32, 14, 32, 630,
	11, 32, 3, 32, 3, 32, 7, 32, 634, 10, 32, 12, 32, 14, 32, 637, 11, 32,
	3, 32, 7, 32, 640, 10, 32, 12, 32, 14, 32, 643, 11, 32, 3, 32, 3, 32, 7,
	32, 647, 10, 32, 12, 32, 14, 32, 650, 11, 32, 5, 32, 652, 10, 32, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 660, 10, 33, 3, 34, 3, 34, 7,
	34, 664, 10, 34, 12, 34, 14, 34, 667, 11, 34, 3, 34, 3, 34, 3, 34, 5, 34,
	672, 10, 34, 3, 34, 7, 34, 675, 10, 34, 12, 34, 14, 34, 678, 11, 34, 3,
	34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	2, 4, 26, 28, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 2, 5, 3, 2, 44, 45, 3, 2, 22, 41, 4, 2, 43, 43, 48, 51, 2, 760,
	2, 72, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6, 96, 3, 2, 2, 2, 8, 139, 3, 2,
	2, 2, 10, 144, 3, 2, 2, 2, 12, 152, 3, 2, 2, 2, 14, 154, 3, 2, 2, 2, 16,
	172, 3, 2, 2, 2, 18, 200, 3, 2, 2, 2, 20, 228, 3, 2, 2, 2, 22, 230, 3,
	2, 2, 2, 24, 238, 3, 2, 2, 2, 26, 240, 3, 2, 2, 2, 28, 307, 3, 2, 2, 2,
	30, 324, 3, 2, 2, 2, 32, 362, 3, 2, 2, 2, 34, 364, 3, 2, 2, 2, 36, 366,
	3, 2, 2, 2, 38, 370, 3, 2, 2, 2, 40, 372, 3, 2, 2, 2, 42, 414, 3, 2, 2,
	2, 44, 416, 3, 2, 2, 2, 46, 424, 3, 2, 2, 2, 48, 426, 3, 2, 2, 2, 50, 498,
	3, 2, 2, 2, 52, 514, 3, 2, 2, 2, 54, 565, 3, 2, 2, 2, 56, 570, 3, 2, 2,
	2, 58, 596, 3, 2, 2, 2, 60, 598, 3, 2, 2, 2, 62, 651, 3, 2, 2, 2, 64, 659,
	3, 2, 2, 2, 66, 661, 3, 2, 2, 2, 68, 682, 3, 2, 2, 2, 70, 687, 3, 2, 2,
	2, 72, 74, 5, 4, 3, 2, 73, 75, 5, 6, 4, 2, 74, 73, 3, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 79, 3, 2, 2, 2, 76, 78, 5, 12, 7, 2, 77, 76, 3, 2, 2, 2,
	78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3,
	2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 7, 2, 2, 3, 83, 3, 3, 2, 2, 2, 84,
	86, 7, 3, 2, 2, 85, 87, 7, 44, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3, 2,
	2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92,
	7, 42, 2, 2, 91, 93, 5, 10, 6, 2, 92, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2,
	2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 5, 3, 2, 2, 2, 96, 97, 7,
	4, 2, 2, 97, 99, 5, 10, 6, 2, 98, 100, 5, 8, 5, 2, 99, 98, 3, 2, 2, 2,
	100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 106,
	3, 2, 2, 2, 103, 105, 5, 10, 6, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2,
	2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 7, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 109, 111, 7, 44, 2, 2, 110, 109, 3, 2, 2, 2, 111, 114,
	3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 115, 3, 2,
	2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 42, 2, 2, 116, 140, 5, 10, 6,
	2, 117, 119, 7, 44, 2, 2, 118, 117, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120,
	118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 120,
	3, 2, 2, 2, 123, 127, 7, 42, 2, 2, 124, 126, 7, 44, 2, 2, 125, 124, 3,
	2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2,
	2, 128, 130, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 134, 7, 5, 2, 2, 131,
	133, 7, 44, 2, 2, 132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132,
	3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 134, 3, 2,
	2, 2, 137, 138, 7, 43, 2, 2, 138, 140, 5, 10, 6, 2, 139, 112, 3, 2, 2,
	2, 139, 120, 3, 2, 2, 2, 140, 9, 3, 2, 2, 2, 141, 143, 7, 44, 2, 2, 142,
	141, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145,
	3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 148, 7, 45,
	2, 2, 148, 11, 3, 2, 2, 2, 149, 153, 5, 44, 23, 2, 150, 153, 5, 14, 8,
	2, 151, 153, 5, 16, 9, 2, 152, 149, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152,
	151, 3, 2, 2, 2, 153, 13, 3, 2, 2, 2, 154, 155, 7, 6, 2, 2, 155, 159, 5,
	54, 28, 2, 156, 158, 7, 44, 2, 2, 157, 156, 3, 2, 2, 2, 158, 161, 3, 2,
	2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2,
	161, 159, 3, 2, 2, 2, 162, 164, 7, 45, 2, 2, 163, 162, 3, 2, 2, 2, 163,
	164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 169, 5, 18, 10, 2, 166, 168,
	9, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2,
	2, 2, 169, 170, 3, 2, 2, 2, 170, 15, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2,
	172, 176, 7, 6, 2, 2, 173, 175, 7, 44, 2, 2, 174, 173, 3, 2, 2, 2, 175,
	178, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 179,
	3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 179, 180, 7, 42, 2, 2, 180, 184, 5, 54,
	28, 2, 181, 183, 7, 44, 2, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3, 2, 2,
	2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186,
	184, 3, 2, 2, 2, 187, 189, 7, 45, 2, 2, 188, 187, 3, 2, 2, 2, 188, 189,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 194, 5, 18, 10, 2, 191, 193, 9,
	2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2,
	2, 194, 195, 3, 2, 2, 2, 195, 17, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197,
	199, 7, 44, 2, 2, 198, 197, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198,
	3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 203, 3, 2, 2, 2, 202, 200, 3, 2,
	2, 2, 203, 207, 7, 7, 2, 2, 204, 206, 9, 2, 2, 2, 205, 204, 3, 2, 2, 2,
	206, 209, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	220, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 211, 5, 20, 11, 2, 211, 215,
	5, 10, 6, 2, 212, 214, 9, 2, 2, 2, 213, 212, 3, 2, 2, 2, 214, 217, 3, 2,
	2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 219, 3, 2, 2, 2,
	217, 215, 3, 2, 2, 2, 218, 210, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220,
	218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222, 220,
	3, 2, 2, 2, 223, 224, 7, 8, 2, 2, 224, 19, 3, 2, 2, 2, 225, 229, 5, 24,
	13, 2, 226, 229, 5, 26, 14, 2, 227, 229, 5, 22, 12, 2, 228, 225, 3, 2,
	2, 2, 228, 226, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229, 21, 3, 2, 2, 2,
	230, 232, 7, 9, 2, 2, 231, 233, 7, 44, 2, 2, 232, 231, 3, 2, 2, 2, 233,
	234, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236,
	3, 2, 2, 2, 236, 237, 5, 26, 14, 2, 237, 23, 3, 2, 2, 2, 238, 239, 5, 42,
	22, 2, 239, 25, 3, 2, 2, 2, 240, 241, 8, 14, 1, 2, 241, 242, 5, 28, 15,
	2, 242, 260, 3, 2, 2, 2, 243, 247, 12, 3, 2, 2, 244, 246, 7, 44, 2, 2,
	245, 244, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 250, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 254,
	9, 3, 2, 2, 251, 253, 7, 44, 2, 2, 252, 251, 3, 2, 2, 2, 253, 256, 3, 2,
	2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 257, 3, 2, 2, 2,
	256, 254, 3, 2, 2, 2, 257, 259, 5, 26, 14, 4, 258, 243, 3, 2, 2, 2, 259,
	262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 27, 3,
	2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 264, 8, 15, 1, 2, 264, 308, 5, 32,
	17, 2, 265, 269, 7, 12, 2, 2, 266, 268, 7, 44, 2, 2, 267, 266, 3, 2, 2,
	2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270,
	298, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 295, 5, 28, 15, 2, 273, 275,
	7, 44, 2, 2, 274, 273, 3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274, 3, 2,
	2, 2, 276, 277, 3, 2, 2, 2, 277, 279, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2,
	279, 283, 7, 13, 2, 2, 280, 282, 7, 44, 2, 2, 281, 280, 3, 2, 2, 2, 282,
	285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 286,
	3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 290, 5, 28, 15, 2, 287, 289, 7,
	44, 2, 2, 288, 287, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3, 2, 2,
	2, 290, 291, 3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293,
	276, 3, 2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 295, 296,
	3, 2, 2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 298, 272, 3, 2,
	2, 2, 298, 299, 3, 2, 2, 2, 299, 303, 3, 2, 2, 2, 300, 302, 7, 44, 2, 2,
	301, 300, 3, 2, 2, 2, 302, 305, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303,
	304, 3, 2, 2, 2, 304, 306, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 306, 308,
	7, 14, 2, 2, 307, 263, 3, 2, 2, 2, 307, 265, 3, 2, 2, 2, 308, 321, 3, 2,
	2, 2, 309, 310, 12, 7, 2, 2, 310, 311, 7, 10, 2, 2, 311, 320, 7, 42, 2,
	2, 312, 313, 12, 6, 2, 2, 313, 314, 7, 10, 2, 2, 314, 320, 5, 56, 29, 2,
	315, 316, 12, 5, 2, 2, 316, 320, 7, 11, 2, 2, 317, 318, 12, 4, 2, 2, 318,
	320, 5, 30, 16, 2, 319, 309, 3, 2, 2, 2, 319, 312, 3, 2, 2, 2, 319, 315,
	3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 320, 323, 3, 2, 2, 2, 321, 319, 3, 2,
	2, 2, 321, 322, 3, 2, 2, 2, 322, 29, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2,
	324, 328, 7, 12, 2, 2, 325, 327, 7, 44, 2, 2, 326, 325, 3, 2, 2, 2, 327,
	330, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 333,
	3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 331, 334, 5, 28, 15, 2, 332, 334, 5,
	70, 36, 2, 333, 331, 3, 2, 2, 2, 333, 332, 3, 2, 2, 2, 334, 338, 3, 2,
	2, 2, 335, 337, 7, 44, 2, 2, 336, 335, 3, 2, 2, 2, 337, 340, 3, 2, 2, 2,
	338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 341, 3, 2, 2, 2, 340,
	338, 3, 2, 2, 2, 341, 342, 7, 14, 2, 2, 342, 31, 3, 2, 2, 2, 343, 363,
	5, 56, 29, 2, 344, 363, 5, 38, 20, 2, 345, 349, 7, 15, 2, 2, 346, 348,
	7, 44, 2, 2, 347, 346, 3, 2, 2, 2, 348, 351, 3, 2, 2, 2, 349, 347, 3, 2,
	2, 2, 349, 350, 3, 2, 2, 2, 350, 352, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2,
	352, 356, 5, 26, 14, 2, 353, 355, 7, 44, 2, 2, 354, 353, 3, 2, 2, 2, 355,
	358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 359,
	3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 360, 7, 16, 2, 2, 360, 363, 3, 2,
	2, 2, 361, 363, 5, 34, 18, 2, 362, 343, 3, 2, 2, 2, 362, 344, 3, 2, 2,
	2, 362, 345, 3, 2, 2, 2, 362, 361, 3, 2, 2, 2, 363, 33, 3, 2, 2, 2, 364,
	365, 5, 36, 19, 2, 365, 35, 3, 2, 2, 2, 366, 367, 9, 4, 2, 2, 367, 37,
	3, 2, 2, 2, 368, 371, 7, 42, 2, 2, 369, 371, 5, 40, 21, 2, 370, 368, 3,
	2, 2, 2, 370, 369, 3, 2, 2, 2, 371, 39, 3, 2, 2, 2, 372, 373, 7, 42, 2,
	2, 373, 374, 7, 10, 2, 2, 374, 375, 7, 42, 2, 2, 375, 41, 3, 2, 2, 2, 376,
	378, 7, 17, 2, 2, 377, 379, 7, 44, 2, 2, 378, 377, 3, 2, 2, 2, 379, 380,
	3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 3, 2,
	2, 2, 382, 386, 7, 42, 2, 2, 383, 385, 7, 44, 2, 2, 384, 383, 3, 2, 2,
	2, 385, 388, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387,
	391, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 389, 392, 7, 42, 2, 2, 390, 392,
	5, 56, 29, 2, 391, 389, 3, 2, 2, 2, 391, 390, 3, 2, 2, 2, 392, 415, 3,
	2, 2, 2, 393, 395, 7, 17, 2, 2, 394, 396, 7, 44, 2, 2, 395, 394, 3, 2,
	2, 2, 396, 397, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2,
	398, 399, 3, 2, 2, 2, 399, 403, 7, 42, 2, 2, 400, 402, 7, 44, 2, 2, 401,
	400, 3, 2, 2, 2, 402, 405, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404,
	3, 2, 2, 2, 404, 406, 3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 406, 410, 7, 38,
	2, 2, 407, 409, 7, 44, 2, 2, 408, 407, 3, 2, 2, 2, 409, 412, 3, 2, 2, 2,
	410, 408, 3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 413, 3, 2, 2, 2, 412,
	410, 3, 2, 2, 2, 413, 415, 5, 26, 14, 2, 414, 376, 3, 2, 2, 2, 414, 393,
	3, 2, 2, 2, 415, 43, 3, 2, 2, 2, 416, 418, 5, 46, 24, 2, 417, 419, 5, 10,
	6, 2, 418, 417, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2,
	420, 421, 3, 2, 2, 2, 421, 45, 3, 2, 2, 2, 422, 425, 5, 48, 25, 2, 423,
	425, 5, 52, 27, 2, 424, 422, 3, 2, 2, 2, 424, 423, 3, 2, 2, 2, 425, 47,
	3, 2, 2, 2, 426, 428, 7, 18, 2, 2, 427, 429, 7, 44, 2, 2, 428, 427, 3,
	2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2,
	2, 431, 432, 3, 2, 2, 2, 432, 436, 7, 42, 2, 2, 433, 435, 7, 44, 2, 2,
	434, 433, 3, 2, 2, 2, 435, 438, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436,
	437, 3, 2, 2, 2, 437, 457, 3, 2, 2, 2, 438, 436, 3, 2, 2, 2, 439, 453,
	7, 19, 2, 2, 440, 442, 7, 44, 2, 2, 441, 440, 3, 2, 2, 2, 442, 445, 3,
	2, 2, 2, 443, 441, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 446, 3, 2, 2,
	2, 445, 443, 3, 2, 2, 2, 446, 450, 7, 42, 2, 2, 447, 449, 7, 44, 2, 2,
	448, 447, 3, 2, 2, 2, 449, 452, 3, 2, 2, 2, 450, 448, 3, 2, 2, 2, 450,
	451, 3, 2, 2, 2, 451, 454, 3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 453, 443,
	3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 453, 3, 2, 2, 2, 455, 456, 3, 2,
	2, 2, 456, 458, 3, 2, 2, 2, 457, 439, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2,
	458, 462, 3, 2, 2, 2, 459, 461, 7, 44, 2, 2, 460, 459, 3, 2, 2, 2, 461,
	464, 3, 2, 2, 2, 462, 460, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 466,
	3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 465, 467, 7, 45, 2, 2, 466, 465, 3, 2,
	2, 2, 466, 467, 3, 2, 2, 2, 467, 471, 3, 2, 2, 2, 468, 470, 7, 44, 2, 2,
	469, 468, 3, 2, 2, 2, 470, 473, 3, 2, 2, 2, 471, 469, 3, 2, 2, 2, 471,
	472, 3, 2, 2, 2, 472, 474, 3, 2, 2, 2, 473, 471, 3, 2, 2, 2, 474, 478,
	7, 7, 2, 2, 475, 477, 9, 2, 2, 2, 476, 475, 3, 2, 2, 2, 477, 480, 3, 2,
	2, 2, 478, 476, 3, 2, 2, 2, 478, 479, 3, 2, 2, 2, 479, 484, 3, 2, 2, 2,
	480, 478, 3, 2, 2, 2, 481, 483, 5, 50, 26, 2, 482, 481, 3, 2, 2, 2, 483,
	486, 3, 2, 2, 2, 484, 482, 3, 2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 490,
	3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 487, 489, 9, 2, 2, 2, 488, 487, 3, 2,
	2, 2, 489, 492, 3, 2, 2, 2, 490, 488, 3, 2, 2, 2, 490, 491, 3, 2, 2, 2,
	491, 493, 3, 2, 2, 2, 492, 490, 3, 2, 2, 2, 493, 494, 7, 8, 2, 2, 494,
	49, 3, 2, 2, 2, 495, 497, 7, 44, 2, 2, 496, 495, 3, 2, 2, 2, 497, 500,
	3, 2, 2, 2, 498, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 501, 3, 2,
	2, 2, 500, 498, 3, 2, 2, 2, 501, 505, 7, 42, 2, 2, 502, 504, 7, 44, 2,
	2, 503, 502, 3, 2, 2, 2, 504, 507, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 505,
	506, 3, 2, 2, 2, 506, 508, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 508, 510,
	5, 64, 33, 2, 509, 511, 5, 10, 6, 2, 510, 509, 3, 2, 2, 2, 511, 512, 3,
	2, 2, 2, 512, 510, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 51, 3, 2, 2,
	2, 514, 516, 7, 20, 2, 2, 515, 517, 7, 44, 2, 2, 516, 515, 3, 2, 2, 2,
	517, 518, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2, 518, 519, 3, 2, 2, 2, 519,
	520, 3, 2, 2, 2, 520, 524, 7, 42, 2, 2, 521, 523, 7, 44, 2, 2, 522, 521,
	3, 2, 2, 2, 523, 526, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 524, 525, 3, 2,
	2, 2, 525, 528, 3, 2, 2, 2, 526, 524, 3, 2, 2, 2, 527, 529, 7, 45, 2, 2,
	528, 527, 3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529, 533, 3, 2, 2, 2, 530,
	532, 7, 44, 2, 2, 531, 530, 3, 2, 2, 2, 532, 535, 3, 2, 2, 2, 533, 531,
	3, 2, 2, 2, 533, 534, 3, 2, 2, 2, 534, 536, 3, 2, 2, 2, 535, 533, 3, 2,
	2, 2, 536, 540, 7, 7, 2, 2, 537, 539, 9, 2, 2, 2, 538, 537, 3, 2, 2, 2,
	539, 542, 3, 2, 2, 2, 540, 538, 3, 2, 2, 2, 540, 541, 3, 2, 2, 2, 541,
	551, 3, 2, 2, 2, 542, 540, 3, 2, 2, 2, 543, 545, 5, 54, 28, 2, 544, 546,
	5, 10, 6, 2, 545, 544, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 545, 3, 2,
	2, 2, 547, 548, 3, 2, 2, 2, 548, 550, 3, 2, 2, 2, 549, 543, 3, 2, 2, 2,
	550, 553, 3, 2, 2, 2, 551, 549, 3, 2, 2, 2, 551, 552, 3, 2, 2, 2, 552,
	557, 3, 2, 2, 2, 553, 551, 3, 2, 2, 2, 554, 556, 9, 2, 2, 2, 555, 554,
	3, 2, 2, 2, 556, 559, 3, 2, 2, 2, 557, 555, 3, 2, 2, 2, 557, 558, 3, 2,
	2, 2, 558, 560, 3, 2, 2, 2, 559, 557, 3, 2, 2, 2, 560, 561, 7, 8, 2, 2,
	561, 53, 3, 2, 2, 2, 562, 564, 7, 44, 2, 2, 563, 562, 3, 2, 2, 2, 564,
	567, 3, 2, 2, 2, 565, 563, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 568,
	3, 2, 2, 2, 567, 565, 3, 2, 2, 2, 568, 569, 5, 60, 31, 2, 569, 55, 3, 2,
	2, 2, 570, 571, 7, 42, 2, 2, 571, 575, 7, 15, 2, 2, 572, 574, 7, 44, 2,
	2, 573, 572, 3, 2, 2, 2, 574, 577, 3, 2, 2, 2, 575, 573, 3, 2, 2, 2, 575,
	576, 3, 2, 2, 2, 576, 586, 3, 2, 2, 2, 577, 575, 3, 2, 2, 2, 578, 583,
	5, 58, 30, 2, 579, 580, 7, 13, 2, 2, 580, 582, 5, 58, 30, 2, 581, 579,
	3, 2, 2, 2, 582, 585, 3, 2, 2, 2, 583, 581, 3, 2, 2, 2, 583, 584, 3, 2,
	2, 2, 584, 587, 3, 2, 2, 2, 585, 583, 3, 2, 2, 2, 586, 578, 3, 2, 2, 2,
	586, 587, 3, 2, 2, 2, 587, 591, 3, 2, 2, 2, 588, 590, 7, 44, 2, 2, 589,
	588, 3, 2, 2, 2, 590, 593, 3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 591, 592,
	3, 2, 2, 2, 592, 594, 3, 2, 2, 2, 593, 591, 3, 2, 2, 2, 594, 595, 7, 16,
	2, 2, 595, 57, 3, 2, 2, 2, 596, 597, 5, 26, 14, 2, 597, 59, 3, 2, 2, 2,
	598, 599, 7, 42, 2, 2, 599, 603, 7, 15, 2, 2, 600, 602, 7, 44, 2, 2, 601,
	600, 3, 2, 2, 2, 602, 605, 3, 2, 2, 2, 603, 601, 3, 2, 2, 2, 603, 604,
	3, 2, 2, 2, 604, 614, 3, 2, 2, 2, 605, 603, 3, 2, 2, 2, 606, 611, 5, 62,
	32, 2, 607, 608, 7, 13, 2, 2, 608, 610, 5, 62, 32, 2, 609, 607, 3, 2, 2,
	2, 610, 613, 3, 2, 2, 2, 611, 609, 3, 2, 2, 2, 611, 612, 3, 2, 2, 2, 612,
	615, 3, 2, 2, 2, 613, 611, 3, 2, 2, 2, 614, 606, 3, 2, 2, 2, 614, 615,
	3, 2, 2, 2, 615, 616, 3, 2, 2, 2, 616, 617, 7, 16, 2, 2, 617, 61, 3, 2,
	2, 2, 618, 620, 7, 44, 2, 2, 619, 618, 3, 2, 2, 2, 620, 623, 3, 2, 2, 2,
	621, 619, 3, 2, 2, 2, 621, 622, 3, 2, 2, 2, 622, 624, 3, 2, 2, 2, 623,
	621, 3, 2, 2, 2, 624, 628, 7, 42, 2, 2, 625, 627, 7, 44, 2, 2, 626, 625,
	3, 2, 2, 2, 627, 630, 3, 2, 2, 2, 628, 626, 3, 2, 2, 2, 628, 629, 3, 2,
	2, 2, 629, 631, 3, 2, 2, 2, 630, 628, 3, 2, 2, 2, 631, 635, 5, 64, 33,
	2, 632, 634, 7, 44, 2, 2, 633, 632, 3, 2, 2, 2, 634, 637, 3, 2, 2, 2, 635,
	633, 3, 2, 2, 2, 635, 636, 3, 2, 2, 2, 636, 652, 3, 2, 2, 2, 637, 635,
	3, 2, 2, 2, 638, 640, 7, 44, 2, 2, 639, 638, 3, 2, 2, 2, 640, 643, 3, 2,
	2, 2, 641, 639, 3, 2, 2, 2, 641, 642, 3, 2, 2, 2, 642, 644, 3, 2, 2, 2,
	643, 641, 3, 2, 2, 2, 644, 648, 5, 64, 33, 2, 645, 647, 7, 44, 2, 2, 646,
	645, 3, 2, 2, 2, 647, 650, 3, 2, 2, 2, 648, 646, 3, 2, 2, 2, 648, 649,
	3, 2, 2, 2, 649, 652, 3, 2, 2, 2, 650, 648, 3, 2, 2, 2, 651, 621, 3, 2,
	2, 2, 651, 641, 3, 2, 2, 2, 652, 63, 3, 2, 2, 2, 653, 660, 7, 42, 2, 2,
	654, 655, 7, 42, 2, 2, 655, 656, 7, 10, 2, 2, 656, 660, 7, 42, 2, 2, 657,
	660, 5, 66, 34, 2, 658, 660, 5, 68, 35, 2, 659, 653, 3, 2, 2, 2, 659, 654,
	3, 2, 2, 2, 659, 657, 3, 2, 2, 2, 659, 658, 3, 2, 2, 2, 660, 65, 3, 2,
	2, 2, 661, 665, 7, 12, 2, 2, 662, 664, 7, 44, 2, 2, 663, 662, 3, 2, 2,
	2, 664, 667, 3, 2, 2, 2, 665, 663, 3, 2, 2, 2, 665, 666, 3, 2, 2, 2, 666,
	671, 3, 2, 2, 2, 667, 665, 3, 2, 2, 2, 668, 672, 7, 42, 2, 2, 669, 672,
	7, 48, 2, 2, 670, 672, 5, 70, 36, 2, 671, 668, 3, 2, 2, 2, 671, 669, 3,
	2, 2, 2, 671, 670, 3, 2, 2, 2, 671, 672, 3, 2, 2, 2, 672, 676, 3, 2, 2,
	2, 673, 675, 7, 44, 2, 2, 674, 673, 3, 2, 2, 2, 675, 678, 3, 2, 2, 2, 676,
	674, 3, 2, 2, 2, 676, 677, 3, 2, 2, 2, 677, 679, 3, 2, 2, 2, 678, 676,
	3, 2, 2, 2, 679, 680, 7, 14, 2, 2, 680, 681, 7, 42, 2, 2, 681, 67, 3, 2,
	2, 2, 682, 683, 7, 21, 2, 2, 683, 684, 5, 64, 33, 2, 684, 685, 7, 14, 2,
	2, 685, 686, 5, 64, 33, 2, 686, 69, 3, 2, 2, 2, 687, 688, 7, 35, 2, 2,
	688, 71, 3, 2, 2, 2, 98, 74, 79, 88, 94, 101, 106, 112, 120, 127, 134,
	139, 144, 152, 159, 163, 169, 176, 184, 188, 194, 200, 207, 215, 220, 228,
	234, 247, 254, 260, 269, 276, 283, 290, 295, 298, 303, 307, 319, 321, 328,
	333, 338, 349, 356, 362, 370, 380, 386, 391, 397, 403, 410, 414, 420, 424,
	430, 436, 443, 450, 455, 457, 462, 466, 471, 478, 484, 490, 498, 505, 512,
	518, 524, 528, 533, 540, 547, 551, 557, 565, 575, 583, 586, 591, 603, 611,
	614, 621, 628, 635, 641, 648, 651, 659, 665, 671, 676,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'module'", "'imports:'", "':'", "'func'", "'{'", "'}'", "'return'",
	"'.'", "'...'", "'['", "','", "']'", "'('", "')'", "'var'", "'type'", "'implements'",
	"'interface'", "'map['", "'||'", "'&&'", "'=='", "'!='", "'<'", "'>'",
	"'<='", "'>='", "'+'", "'-'", "'<<'", "'>>'", "'^'", "'*'", "'/'", "'%'",
	"'='", "'|'", "'&'", "'&^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "OR", "AND", "EQ", "NE", "LT", "GT", "LE", "GE", "PLUS", "MINUS",
	"LSHIFT", "RSHIFT", "CARET", "MUL", "DIV", "REM", "ASSIGN", "BITWISE_OR",
	"BITWISE_AND", "AND_NOT", "IDENTIFIER", "STRING_LIT", "WS", "NEWLINE",
	"LITTLE_U_VALUE", "BIG_U_VALUE", "INT_LIT", "FLOAT_LIT", "IMAGINARY_LIT",
	"RUNE_LIT",
}

var ruleNames = []string{
	"prog", "packageClause", "importDecl", "importSpec", "eos", "topLevelDecl",
	"funcDecl", "methodDecl", "block", "statement", "returnExpr", "statementDecl",
	"expression", "primaryExpr", "arraySelection", "operand", "literal", "basicLit",
	"operandName", "qualifiedIdent", "varDecl", "declaration", "typeDecl",
	"structDecl", "typeSpec", "interfaceDecl", "methodSpec", "funcCallSpec",
	"funcCallArg", "funcSpec", "funcArg", "typeRule", "arrayType", "mapType",
	"star",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ConcertoParser struct {
	*antlr.BaseParser
}

func NewConcertoParser(input antlr.TokenStream) *ConcertoParser {
	this := new(ConcertoParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Concerto.G4"

	return this
}

// ConcertoParser tokens.
const (
	ConcertoParserEOF            = antlr.TokenEOF
	ConcertoParserT__0           = 1
	ConcertoParserT__1           = 2
	ConcertoParserT__2           = 3
	ConcertoParserT__3           = 4
	ConcertoParserT__4           = 5
	ConcertoParserT__5           = 6
	ConcertoParserT__6           = 7
	ConcertoParserT__7           = 8
	ConcertoParserT__8           = 9
	ConcertoParserT__9           = 10
	ConcertoParserT__10          = 11
	ConcertoParserT__11          = 12
	ConcertoParserT__12          = 13
	ConcertoParserT__13          = 14
	ConcertoParserT__14          = 15
	ConcertoParserT__15          = 16
	ConcertoParserT__16          = 17
	ConcertoParserT__17          = 18
	ConcertoParserT__18          = 19
	ConcertoParserOR             = 20
	ConcertoParserAND            = 21
	ConcertoParserEQ             = 22
	ConcertoParserNE             = 23
	ConcertoParserLT             = 24
	ConcertoParserGT             = 25
	ConcertoParserLE             = 26
	ConcertoParserGE             = 27
	ConcertoParserPLUS           = 28
	ConcertoParserMINUS          = 29
	ConcertoParserLSHIFT         = 30
	ConcertoParserRSHIFT         = 31
	ConcertoParserCARET          = 32
	ConcertoParserMUL            = 33
	ConcertoParserDIV            = 34
	ConcertoParserREM            = 35
	ConcertoParserASSIGN         = 36
	ConcertoParserBITWISE_OR     = 37
	ConcertoParserBITWISE_AND    = 38
	ConcertoParserAND_NOT        = 39
	ConcertoParserIDENTIFIER     = 40
	ConcertoParserSTRING_LIT     = 41
	ConcertoParserWS             = 42
	ConcertoParserNEWLINE        = 43
	ConcertoParserLITTLE_U_VALUE = 44
	ConcertoParserBIG_U_VALUE    = 45
	ConcertoParserINT_LIT        = 46
	ConcertoParserFLOAT_LIT      = 47
	ConcertoParserIMAGINARY_LIT  = 48
	ConcertoParserRUNE_LIT       = 49
)

// ConcertoParser rules.
const (
	ConcertoParserRULE_prog           = 0
	ConcertoParserRULE_packageClause  = 1
	ConcertoParserRULE_importDecl     = 2
	ConcertoParserRULE_importSpec     = 3
	ConcertoParserRULE_eos            = 4
	ConcertoParserRULE_topLevelDecl   = 5
	ConcertoParserRULE_funcDecl       = 6
	ConcertoParserRULE_methodDecl     = 7
	ConcertoParserRULE_block          = 8
	ConcertoParserRULE_statement      = 9
	ConcertoParserRULE_returnExpr     = 10
	ConcertoParserRULE_statementDecl  = 11
	ConcertoParserRULE_expression     = 12
	ConcertoParserRULE_primaryExpr    = 13
	ConcertoParserRULE_arraySelection = 14
	ConcertoParserRULE_operand        = 15
	ConcertoParserRULE_literal        = 16
	ConcertoParserRULE_basicLit       = 17
	ConcertoParserRULE_operandName    = 18
	ConcertoParserRULE_qualifiedIdent = 19
	ConcertoParserRULE_varDecl        = 20
	ConcertoParserRULE_declaration    = 21
	ConcertoParserRULE_typeDecl       = 22
	ConcertoParserRULE_structDecl     = 23
	ConcertoParserRULE_typeSpec       = 24
	ConcertoParserRULE_interfaceDecl  = 25
	ConcertoParserRULE_methodSpec     = 26
	ConcertoParserRULE_funcCallSpec   = 27
	ConcertoParserRULE_funcCallArg    = 28
	ConcertoParserRULE_funcSpec       = 29
	ConcertoParserRULE_funcArg        = 30
	ConcertoParserRULE_typeRule       = 31
	ConcertoParserRULE_arrayType      = 32
	ConcertoParserRULE_mapType        = 33
	ConcertoParserRULE_star           = 34
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) PackageClause() IPackageClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageClauseContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConcertoParserEOF, 0)
}

func (s *ProgContext) ImportDecl() IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *ProgContext) AllTopLevelDecl() []ITopLevelDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem())
	var tst = make([]ITopLevelDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDeclContext)
		}
	}

	return tst
}

func (s *ProgContext) TopLevelDecl(i int) ITopLevelDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDeclContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConcertoParserRULE_prog)
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
		p.SetState(70)
		p.PackageClause()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__1 {
		{
			p.SetState(71)
			p.ImportDecl()
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConcertoParserT__3)|(1<<ConcertoParserT__15)|(1<<ConcertoParserT__17))) != 0 {
		{
			p.SetState(74)
			p.TopLevelDecl()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(ConcertoParserEOF)
	}

	return localctx
}

// IPackageClauseContext is an interface to support dynamic dispatch.
type IPackageClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageClauseContext differentiates from other interfaces.
	IsPackageClauseContext()
}

type PackageClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageClauseContext() *PackageClauseContext {
	var p = new(PackageClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_packageClause
	return p
}

func (*PackageClauseContext) IsPackageClauseContext() {}

func NewPackageClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageClauseContext {
	var p = new(PackageClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_packageClause

	return p
}

func (s *PackageClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *PackageClauseContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *PackageClauseContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *PackageClauseContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *PackageClauseContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *PackageClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterPackageClause(s)
	}
}

func (s *PackageClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitPackageClause(s)
	}
}

func (s *PackageClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitPackageClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) PackageClause() (localctx IPackageClauseContext) {
	localctx = NewPackageClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConcertoParserRULE_packageClause)
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
		p.SetState(82)
		p.Match(ConcertoParserT__0)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(83)
			p.Match(ConcertoParserWS)
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(89)
			p.Eos()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportSpecContext)(nil)).Elem())
	var tst = make([]IImportSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportSpecContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConcertoParserRULE_importDecl)
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
		p.SetState(94)
		p.Match(ConcertoParserT__1)
	}
	{
		p.SetState(95)
		p.Eos()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(96)
				p.ImportSpec()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(101)
			p.Eos()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *ImportSpecContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ImportSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ImportSpecContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserSTRING_LIT, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConcertoParserRULE_importSpec)
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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(107)
				p.Match(ConcertoParserWS)
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(113)
			p.Match(ConcertoParserIDENTIFIER)
		}
		{
			p.SetState(114)
			p.Eos()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(115)
				p.Match(ConcertoParserWS)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(121)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(122)
				p.Match(ConcertoParserWS)
			}

			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(128)
			p.Match(ConcertoParserT__2)
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(129)
				p.Match(ConcertoParserWS)
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(135)
			p.Match(ConcertoParserSTRING_LIT)
		}
		{
			p.SetState(136)
			p.Eos()
		}

	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, 0)
}

func (s *EosContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *EosContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConcertoParserRULE_eos)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(139)
			p.Match(ConcertoParserWS)
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.Match(ConcertoParserNEWLINE)
	}

	return localctx
}

// ITopLevelDeclContext is an interface to support dynamic dispatch.
type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}

type TopLevelDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_topLevelDecl
	return p
}

func (*TopLevelDeclContext) IsTopLevelDeclContext() {}

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_topLevelDecl

	return p
}

func (s *TopLevelDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDeclContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TopLevelDeclContext) FuncDecl() IFuncDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopLevelDeclContext) MethodDecl() IMethodDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodDeclContext)
}

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTopLevelDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TopLevelDecl() (localctx ITopLevelDeclContext) {
	localctx = NewTopLevelDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConcertoParserRULE_topLevelDecl)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.FuncDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.MethodDecl()
		}

	}

	return localctx
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcDecl
	return p
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) MethodSpec() IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *FuncDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConcertoParserRULE_funcDecl)
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
		p.SetState(152)
		p.Match(ConcertoParserT__3)
	}
	{
		p.SetState(153)
		p.MethodSpec()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(154)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(160)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	{
		p.SetState(163)
		p.Block()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(164)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodDeclContext is an interface to support dynamic dispatch.
type IMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodDeclContext differentiates from other interfaces.
	IsMethodDeclContext()
}

type MethodDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDeclContext() *MethodDeclContext {
	var p = new(MethodDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_methodDecl
	return p
}

func (*MethodDeclContext) IsMethodDeclContext() {}

func NewMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclContext {
	var p = new(MethodDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_methodDecl

	return p
}

func (s *MethodDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *MethodDeclContext) MethodSpec() IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *MethodDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MethodDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *MethodDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *MethodDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *MethodDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *MethodDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterMethodDecl(s)
	}
}

func (s *MethodDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitMethodDecl(s)
	}
}

func (s *MethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitMethodDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) MethodDecl() (localctx IMethodDeclContext) {
	localctx = NewMethodDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConcertoParserRULE_methodDecl)
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
		p.SetState(170)
		p.Match(ConcertoParserT__3)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(171)
			p.Match(ConcertoParserWS)
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(178)
		p.MethodSpec()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(179)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(185)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	{
		p.SetState(188)
		p.Block()
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(189)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *BlockContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *BlockContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConcertoParserRULE_block)
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
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(195)
			p.Match(ConcertoParserWS)
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(ConcertoParserT__4)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(202)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConcertoParserT__6)|(1<<ConcertoParserT__9)|(1<<ConcertoParserT__12)|(1<<ConcertoParserT__14))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ConcertoParserIDENTIFIER-40))|(1<<(ConcertoParserSTRING_LIT-40))|(1<<(ConcertoParserINT_LIT-40))|(1<<(ConcertoParserFLOAT_LIT-40))|(1<<(ConcertoParserIMAGINARY_LIT-40))|(1<<(ConcertoParserRUNE_LIT-40)))) != 0) {
		{
			p.SetState(208)
			p.Statement()
		}
		{
			p.SetState(209)
			p.Eos()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
			{
				p.SetState(210)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		p.Match(ConcertoParserT__5)
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
	p.RuleIndex = ConcertoParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StatementDecl() IStatementDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementDeclContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) ReturnExpr() IReturnExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnExprContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ConcertoParserRULE_statement)

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

	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.StatementDecl()
		}

	case ConcertoParserT__9, ConcertoParserT__12, ConcertoParserIDENTIFIER, ConcertoParserSTRING_LIT, ConcertoParserINT_LIT, ConcertoParserFLOAT_LIT, ConcertoParserIMAGINARY_LIT, ConcertoParserRUNE_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.expression(0)
		}

	case ConcertoParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.ReturnExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReturnExprContext is an interface to support dynamic dispatch.
type IReturnExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnExprContext differentiates from other interfaces.
	IsReturnExprContext()
}

type ReturnExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnExprContext() *ReturnExprContext {
	var p = new(ReturnExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_returnExpr
	return p
}

func (*ReturnExprContext) IsReturnExprContext() {}

func NewReturnExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnExprContext {
	var p = new(ReturnExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_returnExpr

	return p
}

func (s *ReturnExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ReturnExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ReturnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterReturnExpr(s)
	}
}

func (s *ReturnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitReturnExpr(s)
	}
}

func (s *ReturnExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitReturnExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ReturnExpr() (localctx IReturnExprContext) {
	localctx = NewReturnExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ConcertoParserRULE_returnExpr)
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
		p.SetState(228)
		p.Match(ConcertoParserT__6)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(229)
			p.Match(ConcertoParserWS)
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(234)
		p.expression(0)
	}

	return localctx
}

// IStatementDeclContext is an interface to support dynamic dispatch.
type IStatementDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementDeclContext differentiates from other interfaces.
	IsStatementDeclContext()
}

type StatementDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementDeclContext() *StatementDeclContext {
	var p = new(StatementDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_statementDecl
	return p
}

func (*StatementDeclContext) IsStatementDeclContext() {}

func NewStatementDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementDeclContext {
	var p = new(StatementDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_statementDecl

	return p
}

func (s *StatementDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementDeclContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStatementDecl(s)
	}
}

func (s *StatementDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStatementDecl(s)
	}
}

func (s *StatementDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStatementDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) StatementDecl() (localctx IStatementDeclContext) {
	localctx = NewStatementDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ConcertoParserRULE_statementDecl)

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
		p.SetState(236)
		p.VarDecl()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(ConcertoParserOR, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(ConcertoParserAND, 0)
}

func (s *ExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(ConcertoParserEQ, 0)
}

func (s *ExpressionContext) NE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNE, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserLT, 0)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserLE, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserGT, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserGE, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ConcertoParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ConcertoParserMINUS, 0)
}

func (s *ExpressionContext) BITWISE_OR() antlr.TerminalNode {
	return s.GetToken(ConcertoParserBITWISE_OR, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(ConcertoParserCARET, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(ConcertoParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(ConcertoParserDIV, 0)
}

func (s *ExpressionContext) REM() antlr.TerminalNode {
	return s.GetToken(ConcertoParserREM, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserRSHIFT, 0)
}

func (s *ExpressionContext) BITWISE_AND() antlr.TerminalNode {
	return s.GetToken(ConcertoParserBITWISE_AND, 0)
}

func (s *ExpressionContext) AND_NOT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserAND_NOT, 0)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ConcertoParserASSIGN, 0)
}

func (s *ExpressionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ExpressionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ConcertoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ConcertoParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(239)
		p.primaryExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_expression)
			p.SetState(241)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ConcertoParserWS {
				{
					p.SetState(242)
					p.Match(ConcertoParserWS)
				}

				p.SetState(247)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(248)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ConcertoParserOR-20))|(1<<(ConcertoParserAND-20))|(1<<(ConcertoParserEQ-20))|(1<<(ConcertoParserNE-20))|(1<<(ConcertoParserLT-20))|(1<<(ConcertoParserGT-20))|(1<<(ConcertoParserLE-20))|(1<<(ConcertoParserGE-20))|(1<<(ConcertoParserPLUS-20))|(1<<(ConcertoParserMINUS-20))|(1<<(ConcertoParserLSHIFT-20))|(1<<(ConcertoParserRSHIFT-20))|(1<<(ConcertoParserCARET-20))|(1<<(ConcertoParserMUL-20))|(1<<(ConcertoParserDIV-20))|(1<<(ConcertoParserREM-20))|(1<<(ConcertoParserASSIGN-20))|(1<<(ConcertoParserBITWISE_OR-20))|(1<<(ConcertoParserBITWISE_AND-20))|(1<<(ConcertoParserAND_NOT-20)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ConcertoParserWS {
				{
					p.SetState(249)
					p.Match(ConcertoParserWS)
				}

				p.SetState(254)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(255)
				p.expression(2)
			}

		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *PrimaryExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *PrimaryExprContext) AllPrimaryExpr() []IPrimaryExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem())
	var tst = make([]IPrimaryExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExprContext)
		}
	}

	return tst
}

func (s *PrimaryExprContext) PrimaryExpr(i int) IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) FuncCallSpec() IFuncCallSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallSpecContext)
}

func (s *PrimaryExprContext) ArraySelection() IArraySelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraySelectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraySelectionContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *ConcertoParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ConcertoParserRULE_primaryExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(305)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__12, ConcertoParserIDENTIFIER, ConcertoParserSTRING_LIT, ConcertoParserINT_LIT, ConcertoParserFLOAT_LIT, ConcertoParserIMAGINARY_LIT, ConcertoParserRUNE_LIT:
		{
			p.SetState(262)
			p.Operand()
		}

	case ConcertoParserT__9:
		{
			p.SetState(263)
			p.Match(ConcertoParserT__9)
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(264)
					p.Match(ConcertoParserWS)
				}

			}
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ConcertoParserT__9 || _la == ConcertoParserT__12 || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ConcertoParserIDENTIFIER-40))|(1<<(ConcertoParserSTRING_LIT-40))|(1<<(ConcertoParserINT_LIT-40))|(1<<(ConcertoParserFLOAT_LIT-40))|(1<<(ConcertoParserIMAGINARY_LIT-40))|(1<<(ConcertoParserRUNE_LIT-40)))) != 0) {
			{
				p.SetState(270)
				p.primaryExpr(0)
			}
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					p.SetState(274)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == ConcertoParserWS {
						{
							p.SetState(271)
							p.Match(ConcertoParserWS)
						}

						p.SetState(276)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}
					{
						p.SetState(277)
						p.Match(ConcertoParserT__10)
					}
					p.SetState(281)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == ConcertoParserWS {
						{
							p.SetState(278)
							p.Match(ConcertoParserWS)
						}

						p.SetState(283)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}
					{
						p.SetState(284)
						p.primaryExpr(0)
					}
					p.SetState(288)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

					for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
						if _alt == 1 {
							{
								p.SetState(285)
								p.Match(ConcertoParserWS)
							}

						}
						p.SetState(290)
						p.GetErrorHandler().Sync(p)
						_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
					}

				}
				p.SetState(295)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
			}

		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(298)
				p.Match(ConcertoParserWS)
			}

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(304)
			p.Match(ConcertoParserT__11)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(308)
					p.Match(ConcertoParserT__7)
				}
				{
					p.SetState(309)
					p.Match(ConcertoParserIDENTIFIER)
				}

			case 2:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(311)
					p.Match(ConcertoParserT__7)
				}
				{
					p.SetState(312)
					p.FuncCallSpec()
				}

			case 3:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(314)
					p.Match(ConcertoParserT__8)
				}

			case 4:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(316)
					p.ArraySelection()
				}

			}

		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}

	return localctx
}

// IArraySelectionContext is an interface to support dynamic dispatch.
type IArraySelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArraySelectionContext differentiates from other interfaces.
	IsArraySelectionContext()
}

type ArraySelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArraySelectionContext() *ArraySelectionContext {
	var p = new(ArraySelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_arraySelection
	return p
}

func (*ArraySelectionContext) IsArraySelectionContext() {}

func NewArraySelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraySelectionContext {
	var p = new(ArraySelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_arraySelection

	return p
}

func (s *ArraySelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraySelectionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ArraySelectionContext) Star() IStarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *ArraySelectionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ArraySelectionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ArraySelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraySelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterArraySelection(s)
	}
}

func (s *ArraySelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitArraySelection(s)
	}
}

func (s *ArraySelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitArraySelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ArraySelection() (localctx IArraySelectionContext) {
	localctx = NewArraySelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ConcertoParserRULE_arraySelection)
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
		p.SetState(322)
		p.Match(ConcertoParserT__9)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(323)
			p.Match(ConcertoParserWS)
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__9, ConcertoParserT__12, ConcertoParserIDENTIFIER, ConcertoParserSTRING_LIT, ConcertoParserINT_LIT, ConcertoParserFLOAT_LIT, ConcertoParserIMAGINARY_LIT, ConcertoParserRUNE_LIT:
		{
			p.SetState(329)
			p.primaryExpr(0)
		}

	case ConcertoParserMUL:
		{
			p.SetState(330)
			p.Star()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(333)
			p.Match(ConcertoParserWS)
		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(339)
		p.Match(ConcertoParserT__11)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) FuncCallSpec() IFuncCallSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallSpecContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *OperandContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *OperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ConcertoParserRULE_operand)
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

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.FuncCallSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.OperandName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(343)
			p.Match(ConcertoParserT__12)
		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(344)
				p.Match(ConcertoParserWS)
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(350)
			p.expression(0)
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(351)
				p.Match(ConcertoParserWS)
			}

			p.SetState(356)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(357)
			p.Match(ConcertoParserT__13)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(359)
			p.Literal()
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ConcertoParserRULE_literal)

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
		p.SetState(362)
		p.BasicLit()
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserINT_LIT, 0)
}

func (s *BasicLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserFLOAT_LIT, 0)
}

func (s *BasicLitContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIMAGINARY_LIT, 0)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserRUNE_LIT, 0)
}

func (s *BasicLitContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserSTRING_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterBasicLit(s)
	}
}

func (s *BasicLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitBasicLit(s)
	}
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ConcertoParserRULE_basicLit)
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
		p.SetState(364)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ConcertoParserSTRING_LIT-41))|(1<<(ConcertoParserINT_LIT-41))|(1<<(ConcertoParserFLOAT_LIT-41))|(1<<(ConcertoParserIMAGINARY_LIT-41))|(1<<(ConcertoParserRUNE_LIT-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *OperandNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ConcertoParserRULE_operandName)

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

	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.QualifiedIdent()
		}

	}

	return localctx
}

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_qualifiedIdent
	return p
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitQualifiedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) QualifiedIdent() (localctx IQualifiedIdentContext) {
	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ConcertoParserRULE_qualifiedIdent)

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
		p.SetState(370)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(371)
		p.Match(ConcertoParserT__7)
	}
	{
		p.SetState(372)
		p.Match(ConcertoParserIDENTIFIER)
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *VarDeclContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *VarDeclContext) FuncCallSpec() IFuncCallSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallSpecContext)
}

func (s *VarDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *VarDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *VarDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ConcertoParserRULE_varDecl)
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

	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(ConcertoParserT__14)
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ConcertoParserWS {
			{
				p.SetState(375)
				p.Match(ConcertoParserWS)
			}

			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(380)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(381)
				p.Match(ConcertoParserWS)
			}

			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(387)
				p.Match(ConcertoParserIDENTIFIER)
			}

		case 2:
			{
				p.SetState(388)
				p.FuncCallSpec()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(391)
			p.Match(ConcertoParserT__14)
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ConcertoParserWS {
			{
				p.SetState(392)
				p.Match(ConcertoParserWS)
			}

			p.SetState(395)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(397)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(398)
				p.Match(ConcertoParserWS)
			}

			p.SetState(403)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(404)
			p.Match(ConcertoParserASSIGN)
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(405)
				p.Match(ConcertoParserWS)
			}

			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(411)
			p.expression(0)
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclarationContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *DeclarationContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ConcertoParserRULE_declaration)
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
		p.SetState(414)
		p.TypeDecl()
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(415)
			p.Eos()
		}

		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) StructDecl() IStructDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *TypeDeclContext) InterfaceDecl() IInterfaceDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ConcertoParserRULE_typeDecl)

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

	p.SetState(422)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.StructDecl()
		}

	case ConcertoParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.InterfaceDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_structDecl
	return p
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *StructDeclContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *StructDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *StructDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *StructDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *StructDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *StructDeclContext) AllTypeSpec() []ITypeSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem())
	var tst = make([]ITypeSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecContext)
		}
	}

	return tst
}

func (s *StructDeclContext) TypeSpec(i int) ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ConcertoParserRULE_structDecl)
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
		p.SetState(424)
		p.Match(ConcertoParserT__15)
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(425)
			p.Match(ConcertoParserWS)
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(430)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(431)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__16 {
		{
			p.SetState(437)
			p.Match(ConcertoParserT__16)
		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(441)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == ConcertoParserWS {
					{
						p.SetState(438)
						p.Match(ConcertoParserWS)
					}

					p.SetState(443)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(444)
					p.Match(ConcertoParserIDENTIFIER)
				}
				p.SetState(448)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())

				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(445)
							p.Match(ConcertoParserWS)
						}

					}
					p.SetState(450)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(453)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
		}

	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(457)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(463)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(466)
			p.Match(ConcertoParserWS)
		}

		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(472)
		p.Match(ConcertoParserT__4)
	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(473)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())
	}
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(479)
				p.TypeSpec()
			}

		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext())
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(485)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(491)
		p.Match(ConcertoParserT__5)
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *TypeSpecContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *TypeSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *TypeSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *TypeSpecContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *TypeSpecContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ConcertoParserRULE_typeSpec)
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
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(493)
			p.Match(ConcertoParserWS)
		}

		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(499)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(500)
			p.Match(ConcertoParserWS)
		}

		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(506)
		p.TypeRule()
	}
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(507)
				p.Eos()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext())
	}

	return localctx
}

// IInterfaceDeclContext is an interface to support dynamic dispatch.
type IInterfaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceDeclContext differentiates from other interfaces.
	IsInterfaceDeclContext()
}

type InterfaceDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceDeclContext() *InterfaceDeclContext {
	var p = new(InterfaceDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_interfaceDecl
	return p
}

func (*InterfaceDeclContext) IsInterfaceDeclContext() {}

func NewInterfaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclContext {
	var p = new(InterfaceDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_interfaceDecl

	return p
}

func (s *InterfaceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *InterfaceDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *InterfaceDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *InterfaceDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *InterfaceDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *InterfaceDeclContext) AllMethodSpec() []IMethodSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem())
	var tst = make([]IMethodSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodSpecContext)
		}
	}

	return tst
}

func (s *InterfaceDeclContext) MethodSpec(i int) IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *InterfaceDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *InterfaceDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *InterfaceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterInterfaceDecl(s)
	}
}

func (s *InterfaceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitInterfaceDecl(s)
	}
}

func (s *InterfaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitInterfaceDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) InterfaceDecl() (localctx IInterfaceDeclContext) {
	localctx = NewInterfaceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ConcertoParserRULE_interfaceDecl)
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
		p.SetState(512)
		p.Match(ConcertoParserT__17)
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(513)
			p.Match(ConcertoParserWS)
		}

		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(518)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(519)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(525)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(528)
			p.Match(ConcertoParserWS)
		}

		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(534)
		p.Match(ConcertoParserT__4)
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(535)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext())
	}
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(541)
				p.MethodSpec()
			}
			p.SetState(543)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(542)
						p.Eos()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(545)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 75, p.GetParserRuleContext())
			}

		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())
	}
	p.SetState(555)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(552)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(558)
		p.Match(ConcertoParserT__5)
	}

	return localctx
}

// IMethodSpecContext is an interface to support dynamic dispatch.
type IMethodSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodSpecContext differentiates from other interfaces.
	IsMethodSpecContext()
}

type MethodSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodSpecContext() *MethodSpecContext {
	var p = new(MethodSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_methodSpec
	return p
}

func (*MethodSpecContext) IsMethodSpecContext() {}

func NewMethodSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodSpecContext {
	var p = new(MethodSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_methodSpec

	return p
}

func (s *MethodSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodSpecContext) FuncSpec() IFuncSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSpecContext)
}

func (s *MethodSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *MethodSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *MethodSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterMethodSpec(s)
	}
}

func (s *MethodSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitMethodSpec(s)
	}
}

func (s *MethodSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitMethodSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) MethodSpec() (localctx IMethodSpecContext) {
	localctx = NewMethodSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ConcertoParserRULE_methodSpec)
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
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(560)
			p.Match(ConcertoParserWS)
		}

		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(566)
		p.FuncSpec()
	}

	return localctx
}

// IFuncCallSpecContext is an interface to support dynamic dispatch.
type IFuncCallSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallSpecContext differentiates from other interfaces.
	IsFuncCallSpecContext()
}

type FuncCallSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallSpecContext() *FuncCallSpecContext {
	var p = new(FuncCallSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcCallSpec
	return p
}

func (*FuncCallSpecContext) IsFuncCallSpecContext() {}

func NewFuncCallSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallSpecContext {
	var p = new(FuncCallSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcCallSpec

	return p
}

func (s *FuncCallSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *FuncCallSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncCallSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncCallSpecContext) AllFuncCallArg() []IFuncCallArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncCallArgContext)(nil)).Elem())
	var tst = make([]IFuncCallArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncCallArgContext)
		}
	}

	return tst
}

func (s *FuncCallSpecContext) FuncCallArg(i int) IFuncCallArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncCallArgContext)
}

func (s *FuncCallSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncCallSpec(s)
	}
}

func (s *FuncCallSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncCallSpec(s)
	}
}

func (s *FuncCallSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncCallSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncCallSpec() (localctx IFuncCallSpecContext) {
	localctx = NewFuncCallSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ConcertoParserRULE_funcCallSpec)
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
		p.SetState(568)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(569)
		p.Match(ConcertoParserT__12)
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(570)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext())
	}
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__9 || _la == ConcertoParserT__12 || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ConcertoParserIDENTIFIER-40))|(1<<(ConcertoParserSTRING_LIT-40))|(1<<(ConcertoParserINT_LIT-40))|(1<<(ConcertoParserFLOAT_LIT-40))|(1<<(ConcertoParserIMAGINARY_LIT-40))|(1<<(ConcertoParserRUNE_LIT-40)))) != 0) {
		{
			p.SetState(576)
			p.FuncCallArg()
		}
		p.SetState(581)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserT__10 {
			{
				p.SetState(577)
				p.Match(ConcertoParserT__10)
			}
			{
				p.SetState(578)
				p.FuncCallArg()
			}

			p.SetState(583)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(586)
			p.Match(ConcertoParserWS)
		}

		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(592)
		p.Match(ConcertoParserT__13)
	}

	return localctx
}

// IFuncCallArgContext is an interface to support dynamic dispatch.
type IFuncCallArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallArgContext differentiates from other interfaces.
	IsFuncCallArgContext()
}

type FuncCallArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallArgContext() *FuncCallArgContext {
	var p = new(FuncCallArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcCallArg
	return p
}

func (*FuncCallArgContext) IsFuncCallArgContext() {}

func NewFuncCallArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallArgContext {
	var p = new(FuncCallArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcCallArg

	return p
}

func (s *FuncCallArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FuncCallArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncCallArg(s)
	}
}

func (s *FuncCallArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncCallArg(s)
	}
}

func (s *FuncCallArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncCallArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncCallArg() (localctx IFuncCallArgContext) {
	localctx = NewFuncCallArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ConcertoParserRULE_funcCallArg)

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
		p.SetState(594)
		p.expression(0)
	}

	return localctx
}

// IFuncSpecContext is an interface to support dynamic dispatch.
type IFuncSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSpecContext differentiates from other interfaces.
	IsFuncSpecContext()
}

type FuncSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSpecContext() *FuncSpecContext {
	var p = new(FuncSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcSpec
	return p
}

func (*FuncSpecContext) IsFuncSpecContext() {}

func NewFuncSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSpecContext {
	var p = new(FuncSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcSpec

	return p
}

func (s *FuncSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *FuncSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncSpecContext) AllFuncArg() []IFuncArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncArgContext)(nil)).Elem())
	var tst = make([]IFuncArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncArgContext)
		}
	}

	return tst
}

func (s *FuncSpecContext) FuncArg(i int) IFuncArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncArgContext)
}

func (s *FuncSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncSpec(s)
	}
}

func (s *FuncSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncSpec(s)
	}
}

func (s *FuncSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncSpec() (localctx IFuncSpecContext) {
	localctx = NewFuncSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ConcertoParserRULE_funcSpec)
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
		p.SetState(596)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(597)
		p.Match(ConcertoParserT__12)
	}
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(598)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())
	}
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__9 || _la == ConcertoParserT__18 || _la == ConcertoParserIDENTIFIER || _la == ConcertoParserWS {
		{
			p.SetState(604)
			p.FuncArg()
		}
		p.SetState(609)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserT__10 {
			{
				p.SetState(605)
				p.Match(ConcertoParserT__10)
			}
			{
				p.SetState(606)
				p.FuncArg()
			}

			p.SetState(611)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(614)
		p.Match(ConcertoParserT__13)
	}

	return localctx
}

// IFuncArgContext is an interface to support dynamic dispatch.
type IFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgContext differentiates from other interfaces.
	IsFuncArgContext()
}

type FuncArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgContext() *FuncArgContext {
	var p = new(FuncArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcArg
	return p
}

func (*FuncArgContext) IsFuncArgContext() {}

func NewFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgContext {
	var p = new(FuncArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcArg

	return p
}

func (s *FuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *FuncArgContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *FuncArgContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncArgContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncArg() (localctx IFuncArgContext) {
	localctx = NewFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ConcertoParserRULE_funcArg)
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

	p.SetState(649)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 91, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(619)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(616)
				p.Match(ConcertoParserWS)
			}

			p.SetState(621)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(622)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(626)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(623)
				p.Match(ConcertoParserWS)
			}

			p.SetState(628)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(629)
			p.TypeRule()
		}
		p.SetState(633)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(630)
				p.Match(ConcertoParserWS)
			}

			p.SetState(635)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(639)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(636)
				p.Match(ConcertoParserWS)
			}

			p.SetState(641)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(642)
			p.TypeRule()
		}
		p.SetState(646)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(643)
				p.Match(ConcertoParserWS)
			}

			p.SetState(648)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITypeRuleContext is an interface to support dynamic dispatch.
type ITypeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRuleContext differentiates from other interfaces.
	IsTypeRuleContext()
}

type TypeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRuleContext() *TypeRuleContext {
	var p = new(TypeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *TypeRuleContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *TypeRuleContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *TypeRuleContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (s *TypeRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTypeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TypeRule() (localctx ITypeRuleContext) {
	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ConcertoParserRULE_typeRule)

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

	p.SetState(657)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 92, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(651)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(652)
			p.Match(ConcertoParserIDENTIFIER)
		}
		{
			p.SetState(653)
			p.Match(ConcertoParserT__7)
		}
		{
			p.SetState(654)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(655)
			p.ArrayType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(656)
			p.MapType()
		}

	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *ArrayTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *ArrayTypeContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ArrayTypeContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ArrayTypeContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserINT_LIT, 0)
}

func (s *ArrayTypeContext) Star() IStarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ConcertoParserRULE_arrayType)
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
		p.SetState(659)
		p.Match(ConcertoParserT__9)
	}
	p.SetState(663)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(660)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(665)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext())
	}
	p.SetState(669)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserIDENTIFIER:
		{
			p.SetState(666)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case ConcertoParserINT_LIT:
		{
			p.SetState(667)
			p.Match(ConcertoParserINT_LIT)
		}

	case ConcertoParserMUL:
		{
			p.SetState(668)
			p.Star()
		}

	case ConcertoParserT__11, ConcertoParserWS:

	default:
	}
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(671)
			p.Match(ConcertoParserWS)
		}

		p.SetState(676)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(677)
		p.Match(ConcertoParserT__11)
	}
	{
		p.SetState(678)
		p.Match(ConcertoParserIDENTIFIER)
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) AllTypeRule() []ITypeRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem())
	var tst = make([]ITypeRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeRuleContext)
		}
	}

	return tst
}

func (s *MapTypeContext) TypeRule(i int) ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ConcertoParserRULE_mapType)

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
		p.SetState(680)
		p.Match(ConcertoParserT__18)
	}
	{
		p.SetState(681)
		p.TypeRule()
	}
	{
		p.SetState(682)
		p.Match(ConcertoParserT__11)
	}
	{
		p.SetState(683)
		p.TypeRule()
	}

	return localctx
}

// IStarContext is an interface to support dynamic dispatch.
type IStarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStarContext differentiates from other interfaces.
	IsStarContext()
}

type StarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarContext() *StarContext {
	var p = new(StarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_star
	return p
}

func (*StarContext) IsStarContext() {}

func NewStarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarContext {
	var p = new(StarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_star

	return p
}

func (s *StarContext) GetParser() antlr.Parser { return s.parser }
func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStar(s)
	}
}

func (s *StarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStar(s)
	}
}

func (s *StarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Star() (localctx IStarContext) {
	localctx = NewStarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ConcertoParserRULE_star)

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
		p.SetState(685)
		p.Match(ConcertoParserMUL)
	}

	return localctx
}

func (p *ConcertoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 13:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ConcertoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ConcertoParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
