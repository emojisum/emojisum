// Copyright 2018 Stichting Organism
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//! Emojisum 


//External Crates
extern crate blake2_rfc;
extern crate serde;
extern crate serde_json;
#[macro_use] extern crate serde_derive;

//Imports
use std::error::Error;
use std::fs::File;
use std::path::Path;




//Holds main info parsed from standard
#[derive(Deserialize, Debug)]
pub struct Emojisum {
    version: String,
    description: String,
    // these are an ordered list, referened by a byte (each byte of a checksum digest)
    emojiwords: Vec<Vec<String>>
}

// Words are a set of options to represent an emoji.
// Possible options could be the ":colon_notation:" or a "U+26CF" style codepoint.
//pub type Word = String;



impl Emojisum {

    //Pass a emojimapping JSON to start
    pub fn init(file_path: &str) -> Emojisum {
        let json_file_path = Path::new(file_path);
        let json_file = File::open(json_file_path).expect("file not found");
        let deserialized: Emojisum =
        serde_json::from_reader(json_file).expect("error while reading json");

        return deserialized;
    }

    //from_string()
    //from_bytes()
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
