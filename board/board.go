package board

// Board has an unsigned 64-bit integer that stores the board position of
// every white and black piece
type Board struct {
  whitePawns uint64;
  whiteKnights uint64;
  whiteBishops uint64;
  whiteRooks uint64;
  whiteQueens uint64;
  whiteKing uint64;
  
  blackPawns uint64;
  blackKnights uint64;
  blackBishops uint64;
  blackRooks uint64;
  blackQueens uint64;
  blackKing uint64;
}