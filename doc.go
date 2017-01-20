// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	A package to solve Cryptica puzzle games.


	Context

	Cryptica is a mobile game developed by Pixibots. I have no affiliation
	with the company other than being a satisfied customer. Solving a
	puzzle in Cryptica involves moving tiles on a 7x5 board to match a
	desired configuration. All tiles move symultaneously; their movements
	are, however, constrained by the board edges, and possibly also by
	some additional (unmovable) barriers.


	Description
	
	This Go package can search efficiently for solutions of Cryptica and
	Cryptica-like puzzles, in either depth-first or breadth-first fashion.
	Moreover, it can work on boards of arbitrary size (with some caveats).


	History

	I wrote this software in early 2017 as part of a hands-on effort to
	learn Go and its development tools. I would like to thank Wojciech
	Smalinski at Google for introducing me to this weird but powerful
	programming language.
*/
package cryptica
