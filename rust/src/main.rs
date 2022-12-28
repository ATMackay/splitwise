use std::collections::HashMap;
use csv::Reader;
use csv::Error;
use std::time::Instant;


fn read_file(txs: &mut Vec<Transaction>, filename: String) -> Result<(), Error> {
    let mut rdr = Reader::from_path(filename)?;
    let headers = rdr.headers()?;

    let sender: String = headers[0].to_string();
    let s: i32 = sender.parse().unwrap();

    let receiver: String = headers[1].to_string();
    let r: i32 = receiver.parse().unwrap();


    let amount: String = headers[2].to_string();
    let a: i32 = amount.parse().unwrap();

    add(txs, s, r, a);

    for result in rdr.records() {
        let record = result?;

        let sender: String = record[0].to_string();
        let s: i32 = sender.parse().unwrap();

        let receiver: String = record[1].to_string();
        let r: i32 = receiver.parse().unwrap();


        let amount: String = record[2].to_string();
        let a: i32 = amount.parse().unwrap();

        add(txs, s, r, a);
    }
    Ok(())
}

// transaction
#[derive(Debug)]
#[derive(Clone)]
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

	for tx in txs.iter() {
		if tx.sender == tx.receiver {
			continue
		}
        {
            let p = scores.entry(tx.sender).or_insert(0);
            let new_p = *p - tx.amount;
            scores.insert(tx.sender, new_p);
        }
        {
            let r = scores.entry(tx.receiver).or_insert(0);
            let new_r = *r + tx.amount;
            scores.insert(tx.receiver, new_r);
        }
	}

	scores
}

fn is_zero_sum(scores: &HashMap<i32,i32>) -> bool {
    let mut v = 0;
    for (_key, value) in scores {
        v += value
    }
    v == 0
}

// greedy is a recursive algorithm that simplifies debts
// within the group given the current set of outstanding balances.
// The algorithm is as follows
//
// Step 1)
//
//	Work out maximum creditor and debtor minimum and maximum scores
//
//		    mD - minimum score, the largest debtor in the group
//		    mC - maximum score, the largest creditor in the group
//
//	        maxDebtor - index for debtor with score mD
//	        maxCreditor - index for creditor with score mC
//
// Step 2)
//
// if |mD| == |mC| then create a transaction from maxDebtor to maxCreditor of mC and eliminate both maxDebtor and maxCreditor from the group
// if |mD| <  |mC| then create a transaction from maxDebtor to maxCreditor of |mD| and eliminate maxDebtor from the group
// if |mD| >  |mC| then create a transaction from maxDebtor to maxCreditor of mC and eliminate maxCreditor from the group
//
// Step 3)
//
//	Repeat Step 1) - 2) until the group is eliminated and return
//
// WARNING: Does not protect against malicious inputs,
// To avoid an infinite loop use isZeroSum() to check if the input scores are valid
// before executing this function
fn greedy<'a>(scores: &'a mut HashMap<i32,i32>, txs: &'a mut Vec<Transaction>) -> &'a mut Vec<Transaction> {
	
    let size = scores.keys().len();
    if size == 0 {
		return txs
	}

	let (max_creditor, c) = max_entry(scores);

	let (max_debtor, d) = min_entry(scores);

	if c == -d {
		add(txs, max_debtor, max_creditor, c);
		scores.remove(&max_debtor);
		scores.remove(&max_creditor);
	} else if c > -d {
        add(txs, max_debtor, max_creditor, -d);
		scores.remove(&max_debtor);
        let p = scores.entry(max_creditor).or_insert(0);
        let new_p = *p + d;
        scores.insert(max_creditor, new_p);
	} else if c < -d {
        add(txs, max_debtor, max_creditor, c);
		scores.remove(&max_creditor);
        let r = scores.entry(max_debtor).or_insert(0);
        let new_r = *r + c;
        scores.insert(max_debtor, new_r);
	}

    greedy(scores, txs)
}

fn max_entry(scores: &HashMap<i32,i32>) -> (i32, i32) {
	// find first element
    let mut index = 0;
    let mut value = 0;
    for (k, v) in scores {
        index = *k;
        value = *v;
        break
    }
	// obtain max element
    for (k, v) in scores {
		if v > &value {
			index = *k;
            value = *v;
		}
	}
    (index, value)
}

fn min_entry(scores: &HashMap<i32,i32>) -> (i32,i32) {
	// find first element
    let mut index = 0;
    let mut value = 0;
    for (k, v) in scores {
        index = *k;
        value = *v;
        break
    }
	// obtain max element
    for (k, v) in scores {
		if v < &value {
			index = *k;
            value = *v;
		}
	} 
    (index, value)
}

fn simplify_debts(txs: Vec<Transaction>) -> Vec<Transaction> {

    let mut s = scores(txs.to_vec());

    let t = is_zero_sum(&s);
    if !t {
		panic!("invalid scores, must be zero sum");
	}

    let mut d_0 = vec![];
    let debts = greedy(&mut s, &mut d_0);

    return debts.to_vec()
}

// splitwise
fn main() {

    println!("splitwise algorithm in Rust");

    let mut txs = vec![];

    let mut now = Instant::now();
    if let Err(e) = read_file(&mut txs, "../test_data/input.csv".to_string()){
        panic!("{}", e);
    }
    let mut finish = now.elapsed();
    println!("completed csv load execution in {:.2?}", finish);


    now = Instant::now();
    let debts = simplify_debts(txs);
    finish = now.elapsed();
    println!("completed simplify_debts execution in {:.2?}, settlement transactions {}", finish, debts.len());


    println!("simplified debts {:?}", debts);
}

