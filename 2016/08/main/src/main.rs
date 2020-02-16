use std::fs::File;
use std::io::Read;
use std::io::Write;
use std::thread;
use std::time::Duration;

const N_ROWS: usize = 6;
const N_COLS: usize = 50;
const PLAY_ANIMATIONS: bool = false;
const ANIMATIONS_TIMEOUT_MS: u64 = 40;

trait Parsable {
    fn parse(&mut self, line: &str);
}

struct RectangleInstruction {
    n_rows: usize,
    n_cols: usize,
}
impl Parsable for RectangleInstruction {
    fn parse(&mut self, line: &str) {
        let tokens = line.split(' ').collect::<Vec<&str>>();
        if tokens.len() != 2 {
            panic!("");
        }
        let dimensions = tokens[1].split('x').collect::<Vec<&str>>();
        if dimensions.len() != 2 {
            panic!("");
        }
        self.n_rows = dimensions[1].parse().unwrap();
        self.n_cols = dimensions[0].parse().unwrap();
        //println!("rectangle[{}, {}] parsed", self.n_rows, self.n_cols);
    }
}

trait IRotateInstruction {
    fn target(&mut self) -> &mut usize;
    fn shift(&mut self) -> &mut usize;
}
impl<T> Parsable for T
where
    T: IRotateInstruction,
{
    fn parse(&mut self, line: &str) {
        let tokens = line.split(' ').collect::<Vec<&str>>();
        if tokens.len() != 5 {
            panic!("");
        }
        let row_tokens = tokens[2].split('=').collect::<Vec<&str>>();
        if row_tokens.len() != 2 {
            panic!("");
        }
        *self.target() = row_tokens[1].parse().unwrap();
        *self.shift() = tokens[4].parse().unwrap();
    }
}

struct RotateRowInstruction {
    row: usize,
    shift: usize,
}
impl IRotateInstruction for RotateRowInstruction {
    fn target(&mut self) -> &mut usize {
        &mut self.row
    }
    fn shift(&mut self) -> &mut usize {
        &mut self.shift
    }
}

struct RotateColumnInstruction {
    col: usize,
    shift: usize,
}
impl IRotateInstruction for RotateColumnInstruction {
    fn target(&mut self) -> &mut usize {
        &mut self.col
    }
    fn shift(&mut self) -> &mut usize {
        &mut self.shift
    }
}

struct Display {
    grid: Vec<char>,
    n_rows: usize,
    n_cols: usize,
}
impl Display {
    const SET_CHAR: char = '#';
    const UNSET_CHAR: char = '.';

    fn get_grid(&self, i: usize, j: usize) -> char {
        self.grid[self.get_index(i, j)]
    }
    fn set_default_grid(&mut self, i: usize, j: usize) {
        self.grid[j + i * self.n_cols] = Display::SET_CHAR;
    }
    fn set_grid(&mut self, i: usize, j: usize, c: char) {
        self.grid[j + i * self.n_cols] = c;
    }

    fn get_index(&self, i: usize, j: usize) -> usize {
        j + i * self.n_cols
    }
    fn new(n_rows: usize, n_cols: usize) -> Self {
        let dim = n_rows * n_cols;
        let mut ret = Display {
            grid: Vec::<char>::with_capacity(dim),
            n_rows: n_rows,
            n_cols: n_cols,
        };
        for _ in 0..dim {
            ret.grid.push(Display::UNSET_CHAR);
        }
        ret
    }

    fn print(&self) {
        for i in 0..self.n_rows {
            for j in 0..self.n_cols {
                print!("{}", self.get_grid(i, j));
            }
            println!("");
        }
    }
    fn println(&self) {
        self.print();
        println!("");
    }
    fn show(&self) {
        if !PLAY_ANIMATIONS {
            return;
        }

        print!("\x1B[2J");
        self.print();
        thread::sleep(Duration::from_millis(ANIMATIONS_TIMEOUT_MS));
    }
    fn count_set_pixels(&self) -> usize {
        self.grid
            .iter()
            .filter(|&c| *c == Display::SET_CHAR)
            .count()
    }
}

impl ToString for Display {
    fn to_string(&self) -> String {
        let mut ret = String::new();
        for i in 0..self.n_rows {
            for j in 0..self.n_cols {
                ret.push(self.get_grid(i, j));
            }
            ret.push('\n');
        }
        ret
    }
}

trait Runnable {
    fn run(&self, display: &mut Display);
}
impl Runnable for RectangleInstruction {
    fn run(&self, display: &mut Display) {
        for i in 0..self.n_rows {
            for j in 0..self.n_cols {
                display.set_default_grid(i, j);
            }
        }
    }
}

impl Runnable for RotateRowInstruction {
    fn run(&self, display: &mut Display) {
        let mut row_copy = vec![Display::UNSET_CHAR; display.n_cols];

        for j in 0..display.n_cols {
            row_copy[(j + self.shift) % display.n_cols] = display.get_grid(self.row, j);
        }

        for j in 0..display.n_cols {
            display.set_grid(self.row, j, row_copy[j]);
        }
    }
}

impl Runnable for RotateColumnInstruction {
    fn run(&self, display: &mut Display) {
        let mut col_copy = vec![Display::UNSET_CHAR; display.n_rows];

        for i in 0..display.n_rows {
            col_copy[(i + self.shift) % display.n_rows] = display.get_grid(i, self.col);
        }

        for i in 0..display.n_rows {
            display.set_grid(i, self.col, col_copy[i]);
        }
    }
}

struct Instruction {}
impl Instruction {
    fn new(line: &str) -> Box<Runnable> {
        let tokens = line.split(' ').collect::<Vec<&str>>();
        if tokens[0] == "rect" {
            let mut ret = RectangleInstruction {
                n_rows: 0,
                n_cols: 0,
            };
            ret.parse(line);
            return Box::from(ret);
        } else if tokens[1] == "row" {
            let mut ret = RotateRowInstruction { row: 0, shift: 0 };
            ret.parse(line);
            return Box::from(ret);
        } else if tokens[1] == "column" {
            let mut ret = RotateColumnInstruction { col: 0, shift: 0 };
            ret.parse(line);
            return Box::from(ret);
        } else {
            panic!("{}", line);
        }
    }
}

fn main() {
    // let mut display = Display::new(3, 7);
    // display.show();

    // let i = Instruction::new("rect 3x2");
    // i.run(&mut display);
    // display.show();

    // let i = Instruction::new("rotate column x=1 by 1");
    // i.run(&mut display);
    // display.show();

    // let i = Instruction::new("rotate row y=0 by 4");
    // i.run(&mut display);
    // display.show();

    // let i = Instruction::new("rotate column x=1 by 1");
    // i.run(&mut display);
    // display.show();

    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");
    let lines = &contents.split("\n").collect::<Vec<&str>>();

    let mut display = Display::new(N_ROWS, N_COLS);
    display.show();

    for line in lines {
        let instr = Instruction::new(&line);
        instr.run(&mut display);
        display.show();
    }
    let pixel_count = display.count_set_pixels();
    println!("{}", pixel_count);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(pixel_count.to_string().as_bytes())
        .expect("Unable to write");

    display.println();
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(display.to_string().as_bytes())
        .expect("Unable to write");
}
