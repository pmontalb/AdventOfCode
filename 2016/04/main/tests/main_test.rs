extern crate main;

#[cfg(test)]
mod tests {
    #[test]
    fn case_study_1() {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse("aaaaa-bbb-z-y-x-123[abxyz]");
        assert_eq!(r.actual_hash, "abxyz");
        assert_eq!(r.get_hash(), "abxyz");
        assert_eq!(r.is_real(), true);
    }

    #[test]
    fn case_study_2() {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse("a-b-c-d-e-f-g-h-987[abcde]");
        assert_eq!(r.actual_hash, "abcde");
        assert_eq!(r.get_hash(), "abcde");
        assert_eq!(r.is_real(), true);
    }

    #[test]
    fn case_study_3() {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse("not-a-real-room-404[oarel]");
        assert_eq!(r.actual_hash, "oarel");
        assert_eq!(r.get_hash(), "oarel");
        assert_eq!(r.is_real(), true);
    }

    #[test]
    fn case_study_4() {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse("totally-real-room-200[decoy]");
        assert_eq!(r.actual_hash, "decoy");
        assert_eq!(r.get_hash(), "loart");
        assert_eq!(r.is_real(), false);
    }

    #[test]
    fn case_study_5() {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse("qzmt-zixmtkozy-ivhz-343[aaaaa]");
        assert_eq!(r.actual_hash, "aaaaa");
        assert_eq!(r.get_hash(), "zimth");
        assert_eq!(r.is_real(), false);
        assert_eq!(r.decode(), "very encrypted name 343[fffff]");
    }
}
