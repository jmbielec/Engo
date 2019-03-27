package board

// Generates a Board struct with all the pieces in their initial starting
// positions.
func GenerateStartingBoard() Board {
  // White pieces
  wPawns := setBit(0, A2)
  wPawns = setBit(wPawns, B2)
  wPawns = setBit(wPawns, C2)
  wPawns = setBit(wPawns, D2)
  wPawns = setBit(wPawns, E2)
  wPawns = setBit(wPawns, F2)
  wPawns = setBit(wPawns, G2)
  wPawns = setBit(wPawns, H2)
  
  wKnights := setBit(0, B1)
  wKnights = setBit(wKnights, G1)
  
  wBishops := setBit(0, C1)
  wBishops = setBit(wBishops, F1)
  
  wRooks := setBit(0, A1)
  wRooks = setBit(wRooks, H1)
  
  wQueens := setBit(0, D1)
  
  wKing := setBit(0, E1)
  
  
  // Black pieces
  bPawns := setBit(0, A7)
  bPawns = setBit(bPawns, B7)
  bPawns = setBit(bPawns, C7)
  bPawns = setBit(bPawns, D7)
  bPawns = setBit(bPawns, E7)
  bPawns = setBit(bPawns, F7)
  bPawns = setBit(bPawns, G7)
  bPawns = setBit(bPawns, H7)
  
  bKnights := setBit(0, B8)
  bKnights = setBit(bKnights, G8)
  
  bBishops := setBit(0, C8)
  bBishops = setBit(bBishops, F8)
  
  bRooks := setBit(0, A8)
  bRooks = setBit(bRooks, H8)
  
  bQueens := setBit(0, D8)
  
  bKing := setBit(0, E8)

  
  // Board
  board := Board {
    whitePawns: wPawns,
    whiteKnights: wKnights,
    whiteBishops: wBishops,
    whiteRooks: wRooks,
    whiteQueens: wQueens,
    whiteKing: wKing,
    
    blackPawns: bPawns,
    blackKnights: bKnights,
    blackBishops: bBishops,
    blackRooks: bRooks,
    blackQueens: bQueens,
    blackKing: bKing,
  }
  
  return board
}

// Sets the bit at the proper position on a bitboard given a square s
func setBit(bitboard uint64, s Square) uint64 {
  bitboard |= (1 << uint(s))
  return bitboard
}