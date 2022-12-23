use std::collections::HashMap;

// transaction
struct Transaction {
    sender: i32,
    receiver: i32,
    amount: i32,
}

fn add(txs: &mut Vec<Transaction>, s: i32, r: i32, amt: i32) {
    let new_tx = Transaction{
        sender: s,
        receiver: r,
        amount: amt,
    };
    txs.push(new_tx);
}

// Scores returns a map of balances given a set of transactions between parties
//
// Transaction[0] --> Transaction[1] : Transaction[3]
// results in a map
// scores[Transaction[0]] = - Transaction[3]
// scores[Transaction[1]] = + Transaction[3]
//
// e.g.
//
//	1 --> 2 : 10
//
// scores[1] = -10
// scores[2] = 10
fn scores(txs: Vec<Transaction>) -> HashMap<i32,i32> {

    let mut scores = HashMap::<i32, i32>::new();

    // FIXME inefficient method requiring 2 for loops, please fix 
	for tx in txs.iter() {
		if tx.sender == tx.receiver {
			continue
		}
        let p = scores.entry(tx.sender).or_insert(0);
        let new_p = *p + tx.amount;
        scores.insert(tx.sender, new_p);
	}
    for tx in txs.iter() {
		if tx.sender == tx.receiver {
			continue
		}
        let r = scores.entry(tx.receiver).or_insert(0);
        let new_r = *r - tx.amount;
        scores.insert(tx.receiver, new_r);
	}
	return scores
}

fn is_zero_sum(scores: &HashMap<i32,i32>) -> bool {
    let mut v = 0;
    for (_key, value) in scores {
        v += value
    }
    return v == 0
}

fn main() {
    println!("splitwise algorithm written in Rust....TODO");

    let transactions = vec![Transaction{sender: 1, receiver: 2, amount: 10}, Transaction{sender: 2, receiver: 3, amount: 20}, Transaction{sender: 3, receiver: 1, amount: 10}];
    
    let s = scores(transactions);
    
    let t = is_zero_sum(&s);

    if !t {
		panic!("invalid scores, must be zero sum");
	}

    for (key, value) in &s {
        println!("{}: {}", key, value);
    }
}

