**Sudoku Solver (Go)**    
⚡ A blazing-fast command-line Sudoku solver written in Go*  

**🚀 Features**  
- ✔️ Solves 9×9 Sudoku puzzles with backtracking algorithm  
- 🛡️ Input validation for rows, columns, and 3×3 subgrids  
- ✨ Clean output formatting for solved puzzles  
- 🚫 Error handling** for invalid/unsolvable inputs  
- ⚡ Optimized for performance (solves most puzzles in <1ms)  

---

**📦 Installation**  
```sh
# Install globally
go install github.com/yourusername/sudoku-solver@latest

# Or clone and build
git clone https://github.com/yourusername/sudoku-solver
cd sudoku-solver
go build -o sudoku
```

---

**🛠 Usage**    
**Basic Solving**    
```sh
./sudoku "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
```

**Expected Output**   
```text
5 3 4 |6 7 8 |9 1 2  
6 7 2 |1 9 5 |3 4 8  
1 9 8 |3 4 2 |5 6 7  
------+------+------  
8 5 9 |7 6 1 |4 2 3  
4 2 6 |8 5 3 |7 9 1  
7 1 3 |9 2 4 |8 5 6  
------+------+------  
9 6 1 |5 3 7 |2 8 4  
2 8 7 |4 1 9 |6 3 5  
3 4 5 |2 8 6 |1 7 9  
```

 **Error Cases**  
```sh
# Invalid input (duplicate 5 in top-left subgrid)
./sudoku "55..7...." "6..195..." [...] 
# Output: Error
```

---

 **⚙️ How It Works**  
**Input Requirements**  
| Rule                          | Example Valid Input |  
|-------------------------------|---------------------|  
| 9 arguments (1 per row)       | `".96.4..1."`       |  
| 9 chars per argument          | `"1.......7"`       |  
| Digits 1-9 or `.` for empty   | `"..3.5.8.."`       |  

 **Algorithm**  
1. **Validation**  
   - Checks row/column/subgrid conflicts  
   - Verifies input length and characters  

2. **Backtracking Engine**  
   - Recursively tries digits 1-9 in empty cells  
   - Prunes invalid paths early for efficiency  

3. **Output**  
   - Formats solution with grid lines  
   - Returns "Error" for unsolvable puzzles  

---

**🧪 Testing**  
```sh
 Run unit tests
go test -v

Benchmark
go test -bench=.
```

---

 **📜 License**  
MIT © 2024 [Your Name]  

---

**✨ Pro Tips**  
1. **Need visualization?** Pipe output to `column -t` for perfect alignment:  
   ```sh
   ./sudoku "53..7...." [...] | column -t
   ```  
2. **Want harder puzzles?** Try integrating puzzle generation (great next feature!)  

---

