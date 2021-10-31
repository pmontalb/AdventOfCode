const UNSET: i32 = i32::MAX;
const N_BITS: usize = 2;
pub const N_NEEDLES: usize = 2;

#[derive(Debug, Clone)]
pub struct Bot {
    bits: [i32; N_BITS],
    last: usize,
}

impl Default for Bot {
    fn default() -> Bot {
        Bot {
            bits: [UNSET; N_BITS],
            last: 0,
        }
    }
}

impl Bot {
    pub fn check(&self, needles: &[i32; N_NEEDLES]) -> bool {
        let mut j: usize = 0;
        for i in 0..N_BITS {
            if self.bits[i] != needles[j] {
                continue;
            }
            j += 1;
        }
        j == N_NEEDLES
    }

    pub fn low(&self) -> i32 {
        let min_value = self.bits.iter().min();
        match min_value {
            None => panic!(""),
            Some(i) => *i
        }
    }
    pub fn high(&mut self) -> i32 {
        let max_value = self.bits.iter().max();
        match max_value {
            None => panic!(""),
            Some(i) => { assert!(*i != UNSET); return *i; }
        }
    }
    pub fn is_ready(&self) -> bool {
        let first_unset = self.bits.iter().position(|x| *x == UNSET);
        first_unset == None
    }
    pub fn assign(&mut self, bit: i32) {
        assert!(self.bits[self.last] == UNSET);
        self.bits[self.last] = bit;
        self.last += 1;
        if self.last == N_BITS {
            self.last = 0;
        }
    }
}

#[derive(Debug)]
pub struct Instructions<'a> {
    pub needles: &'a [i32; N_NEEDLES],
    pub pendingInstructions: Vec<&'a str>,
    pub bots: Vec<Bot>,
    pub outputs: Vec<i32>,
    pub checkIdx: usize,
}

impl<'a> Instructions<'a> {
    pub fn new(needles: &'a [i32; N_NEEDLES]) -> Self {
        Instructions {
            needles: needles,
            pendingInstructions: Vec::<&str>::new(),
            bots: Vec::<Bot>::new(),
            outputs: Vec::<i32>::new(),
            checkIdx: usize::MAX,
        }
    }

    fn increase_size<T: Clone + std::default::Default>(container: &mut Vec<T>, newSz: usize) {
        if newSz < container.len() {
            return;
        }
        container.resize(newSz + 1, Default::default());
    }

    fn instruction_worker(&mut self, instruction: &'a str) -> bool {
        //println!("Processing instruction {:?}", instruction);

        let tokens = instruction.split(" ").collect::<Vec<&str>>();

        if tokens.iter().position(|x| *x == "value") != None {
            assert_eq!(tokens.len(), 6);
            let value = tokens[1].parse::<i32>().unwrap();
            let botIdx = tokens[5].parse::<usize>().unwrap();
            Self::increase_size(&mut self.bots, botIdx);

            /*println!(
                "Instruction {:?}: assigning value({:?}) to bot({:?})[# {:?}]",
                instruction, value, self.bots[botIdx], botIdx
            );*/
            self.bots[botIdx].assign(value);
            /*println!(
                "Instruction {:?}: [idx({:?}) bot({:?})] after process\n",
                instruction, botIdx, self.bots[botIdx]
            );*/
        } else if tokens.iter().position(|x| *x == "gives") != None {
            assert_eq!(tokens.len(), 12);
            let sourceIdx = tokens[1].parse::<usize>().unwrap();
            Self::increase_size(&mut self.bots, sourceIdx);
            if !self.bots[sourceIdx].is_ready() {
                /*println!(
                    "Instruction {:?}: Bot({:?})[# {:?}] not ready, adding to pending instructions\n",
                   instruction, self.bots[sourceIdx], sourceIdx
                );*/
                return false;
            }

            if self.bots[sourceIdx].check(self.needles) {
                /*println!(
                    "*** FOUND: Bot({:?})[# {:?}] is comparing needles[{:?}] at instrunction {:?}",
                    self.bots[sourceIdx], sourceIdx, self.needles, instruction
                );*/
                self.checkIdx = sourceIdx;
            }

            // process low
            let isLowDestBot = tokens[5] == "bot";
            let lowDestIdx = tokens[6].parse::<usize>().unwrap();
            Self::increase_size(&mut self.bots, lowDestIdx);
            Self::increase_size(&mut self.outputs, lowDestIdx);

            let lowBit = self.bots[sourceIdx].low();
            if isLowDestBot {
                /*println!(
                    "Instruction {:?}: bot({:?})[# {:?}] -> bot({:?})[# {:?}] giving low({:?})",
                    instruction, self.bots[sourceIdx], sourceIdx, self.bots[lowDestIdx], lowDestIdx, lowBit
                );*/
                self.bots[lowDestIdx].assign(lowBit);
            } else {
                /*println!(
                    "Instruction {:?}: bot({:?})[# {:?}] -> output[# {:?}]({:?} -> {:?}) giving low",
                    instruction, self.bots[sourceIdx], sourceIdx, lowDestIdx, self.outputs[lowDestIdx], lowBit
                );*/
                self.outputs[lowDestIdx] = lowBit;
            }

            // process high
            let isHighDestBot = tokens[10] == "bot";
            let highDestIdx = tokens[11].parse::<usize>().unwrap();
            Self::increase_size(&mut self.bots, highDestIdx);
            Self::increase_size(&mut self.outputs, highDestIdx);

            let highBit = self.bots[sourceIdx].high();
            if isHighDestBot {
                /*println!(
                    "Instruction {:?}: bot({:?})[# {:?}] -> bot({:?})[# {:?}] giving high({:?})",
                    instruction, self.bots[sourceIdx], sourceIdx, self.bots[highDestIdx], highDestIdx, highBit
                );*/
                self.bots[highDestIdx].assign(highBit);
            } else {
                /*println!(
                    "Instruction {:?}: bot({:?})[# {:?}] -> output[# {:?}]({:?} -> {:?}) giving high",
                    instruction, self.bots[sourceIdx], sourceIdx, highDestIdx, self.outputs[highDestIdx], highBit
                );*/
                self.outputs[highDestIdx] = highBit;
            }

            // reset source
            self.bots[sourceIdx] = Default::default();
            assert!(!self.bots[sourceIdx].is_ready());

            /*println!(
                "Instruction {:?}: source[idx({:?}) bot({:?})] low[idx({:?}) bot({:?}) out({:?})] high[idx({:?}) bot({:?}) out({:?})]\n",
                instruction, 
                sourceIdx, self.bots[sourceIdx],
                lowDestIdx, self.bots[lowDestIdx], self.outputs[lowDestIdx],
                highDestIdx, self.bots[highDestIdx], self.outputs[highDestIdx],
            );*/
        } else {
            assert!(false);
        }
        return true;
    }

    pub fn process(&mut self, instructions: &Vec<&'a str>) {
        for instruction in instructions.iter() {
            let res = self.instruction_worker(&instruction);
            if !res {
                self.pendingInstructions.push(instruction);
                continue;
            }

            // now process pending instructions
            let mut newPendingInstructions = Vec::<&str>::new();
            let mut pendingInstructions = Vec::<&str>::new();
            for pendingInstruction in &self.pendingInstructions {
                pendingInstructions.push(pendingInstruction);
            }

            for pendingInstruction in pendingInstructions {
                let res = self.instruction_worker(pendingInstruction);
                if !res {
                    //println!("Instruction {:?}: wan't able to process, remaining in the pending instructions", pendingInstruction);
                    newPendingInstructions.push(pendingInstruction);
                    continue;
                }
                //println!("Instruction {:?}: was pending, got processed: numPending({:?})", pendingInstruction, newPendingInstructions.len());
            }
            self.pendingInstructions = newPendingInstructions;
        }

        while !self.pendingInstructions.is_empty() {
            let mut newPendingInstructions = Vec::<&str>::new();
            let mut pendingInstructions = Vec::<&str>::new();
            for pendingInstruction in &self.pendingInstructions {
                pendingInstructions.push(pendingInstruction);
            }

            for pendingInstruction in pendingInstructions {
                let res = self.instruction_worker(pendingInstruction);
                if !res {
                    //println!("Instruction {:?}: wan't able to process, remaining in the pending instructions", pendingInstruction);
                    newPendingInstructions.push(pendingInstruction);
                    continue;
                }
                //println!("Instruction {:?}: was pending, got processed: numPending({:?})", pendingInstruction, newPendingInstructions.len());
            }
            self.pendingInstructions = newPendingInstructions;
        }
    }
}