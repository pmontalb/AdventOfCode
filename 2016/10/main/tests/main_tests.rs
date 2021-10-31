extern crate main;

#[cfg(test)]
mod tests {
    #[test]
    fn case_study_1() {
        let instructionsStr: Vec<&str> = vec![
            "value 5 goes to bot 2",
            "bot 2 gives low to bot 1 and high to bot 0",
            "value 3 goes to bot 1",
            "bot 1 gives low to output 1 and high to bot 0",
            "bot 0 gives low to output 2 and high to output 0",
            "value 2 goes to bot 2",
        ];

        const needles: [i32; main::aoc::N_NEEDLES] = [5, 2];
        let mut instructions = main::aoc::Instructions::new(&needles);
        instructions.process(&instructionsStr);
       
        println!("outputs={:?}", instructions.outputs);
        println!("pendingInstructions={:?}", instructions.pendingInstructions);

        assert_eq!(instructions.bots.len(), 3);
        assert_eq!(instructions.outputs.len(), 3);
        assert_eq!(instructions.pendingInstructions.len(), 0);

        assert_eq!(2, instructions.checkIdx);
        assert_eq!(instructions.outputs[0], 5);
        assert_eq!(instructions.outputs[1], 2);
        assert_eq!(instructions.outputs[2], 3);
    }
}
