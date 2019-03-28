package board

// Generates a Board struct with all the pieces in their initial starting
// positions.
func GenerateStartingBoard() Board {
  // White pieces
  whitePawns := setBit(0, A2)
  whitePawns = setBit(whitePawns, B2)
  whitePawns = setBit(whitePawns, C2)
  whitePawns = setBit(whitePawns, D2)
  whitePawns = setBit(whitePawns, E2)
  whitePawns = setBit(whitePawns, F2)
  whitePawns = setBit(whitePawns, G2)
  whitePawns = setBit(whitePawns, H2)
  
  whiteKnights := setBit(0, B1)
  whiteKnights = setBit(whiteKnights, G1)
  
  whiteBishops := setBit(0, C1)
  whiteBishops = setBit(whiteBishops, F1)
  
  whiteRooks := setBit(0, A1)
  whiteRooks = setBit(whiteRooks, H1)
  
  whiteQueens := setBit(0, D1)
  
  whiteKing := setBit(0, E1)
  
  
  // Black pieces
  blackPawns := setBit(0, A7)
  blackPawns = setBit(blackPawns, B7)
  blackPawns = setBit(blackPawns, C7)
  blackPawns = setBit(blackPawns, D7)
  blackPawns = setBit(blackPawns, E7)
  blackPawns = setBit(blackPawns, F7)
  blackPawns = setBit(blackPawns, G7)
  blackPawns = setBit(blackPawns, H7)
  
  blackKnights := setBit(0, B8)
  blackKnights = setBit(blackKnights, G8)
  
  blackBishops := setBit(0, C8)
  blackBishops = setBit(blackBishops, F8)
  
  blackRooks := setBit(0, A8)
  blackRooks = setBit(blackRooks, H8)
  
  blackQueens := setBit(0, D8)
  
  blackKing := setBit(0, E8)

  
  // Board
  board := Board {
    whitePawns: whitePawns,
    whiteKnights: whiteKnights,
    whiteBishops: whiteBishops,
    whiteRooks: whiteRooks,
    whiteQueens: whiteQueens,
    whiteKing: whiteKing,
    
    blackPawns: blackPawns,
    blackKnights: blackKnights,
    blackBishops: blackBishops,
    blackRooks: blackRooks,
    blackQueens: blackQueens,
    blackKing: blackKing,
  }
  
  return board
}

// Sets the bit at the proper position on a bitboard given a square s
func setBit(bitboard uint64, s Square) uint64 {
  bitboard |= (1 << uint(s))
  return bitboard
}