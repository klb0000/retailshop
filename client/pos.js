const defaultTransaction = {
    'startTime': Date.now(),
    'completeTime': 0,
    'totalCost': null,
    'cashRecieved': null,
    'tRowsStack' : []
}
console.log('debug mode')


const scanSound = new Audio('barcode-sound.mp3')

const dummySound = new Audio('barcode-sound.mp3')
dummySound.volume = 0
setInterval(()=>{
    dummySound.currentTime = 0
    sound.play()
}, 1000)

function newProduct(id, pName, price, discount, taxP) {
    return {
        'id' : id,
        'pName' : pName,
        'discount' : discount,
        'taxP' : taxP,
        'price' : price
    }
}

function isValidProduct(pd) {
    return pd.hasOwnProperty('id') && pd.hasOwnProperty('pName') && 
    pd.hasOwnProperty('discount') && pd.hasOwnProperty('taxP') && pd.hasOwnProperty('price') 
    
}

// make new transaction row
function makeTRow(product, quantity) {
    const priceAfterDiscount = product.price - product.discount
    const totalCostWithoutTax = priceAfterDiscount * quantity
    const totalTaxAmount = calculateTax(totalCostWithoutTax, product.taxP)
    trow = {
        'pid' : product.id,
        'qty' : quantity,
        'pName' : product.pName,
        'uPrice': product.price,
        'discount' : product.discount,
        'taxP' : product.taxP,
        'taxAmount' : totalTaxAmount,
        'cost' : totalCostWithoutTax
    }

    return trow
} 

function calculateTax(taxP, amount) {

    if (taxP > 0) {
        return  Math.round(amount * taxP/100)
    }
    return 0
}

function newTRowWithNewQty(tRow, newQty) {

    const priceAfterDiscount = tRow.uPrice - tRow.discount
    const totalCostWithoutTax = priceAfterDiscount * newQty
    const totalTaxAmount = calculateTax(totalCostWithoutTax, tRow.taxP)
    
    newTrow = {
        'idx' : null,
        'qty' : newQty,
        'pid' : tRow.pid,
        'pName' : tRow.pName,
        'uPrice': tRow.uPrice,
        'discount' : tRow.discount,
        'taxP' : tRow.taxP,
        'taxAmount' : totalTaxAmount,
        'cost' : totalCostWithoutTax
    }
    return newTrow
}

function  calculateGrandTotal(transaction) {

    tRowsStack = transaction.tRowsStack
    const detail = {
        'totalCost': 0,
        'totalTax' : 0,
        'totalItems' :0,
    }

    for (let i = 0; i < tRowsStack.length; i++) {
        detail.totalCost += tRowsStack[i].cost + tRowsStack[i].taxAmount
        detail.totalTax += tRowsStack[i].taxAmount
        detail.totalItems += tRowsStack[i].qty
    }
    return detail
}

function addProductToTransaction(transaction, product, tracker) {   
        
    // if transaction already has product
        if (tracker.has(product.id)) {
            
            const loc = tracker.get(product.id)
            const tRow = defaultTransaction.tRowsStack[loc]
            const oldQty = tRow.qty

            // get new t-row after updating t-row-quantity
            transaction.tRowsStack[loc] = newTRowWithNewQty(tRow, oldQty+1)
            return
        }
    
        const tRow = makeTRow(product, 1)
        loc = transaction.tRowsStack.length
        tracker.set(product.id, loc)
        transaction.tRowsStack.push(tRow)
}

function popFromStack(stack, loc) {
    stack.splice(loc, 1)
    return stack
}

function updateTracker(tracker, stack) {
    tracker.clear()
    for (let i = 0; i < stack.length; i++ ){
        const tRow = stack[i]
        tracker.set(tRow.pid, i)
    }
}

function removeProductFromTransaction(transaction, productId, tracker) {
    if (tracker.has(productId)) {
        const delSound = new Audio('delete-sound.mp3')
        const loc = tracker.get(productId)
        const tRow = transaction.tRowsStack[loc]
        delSound.play()
        if (tRow.qty > 1) {
        transaction.tRowsStack[loc] = newTRowWithNewQty(tRow, tRow.qty-1)
            return
        }
        transaction.tRowsStack = popFromStack(transaction.tRowsStack, loc)
        updateTracker(tracker, transaction.tRowsStack)
    }
}

// called when delete button is pressed 
function deleteFromTransaction() {
    const productId = document.getElementById("in").value
    removeProductFromTransaction(defaultTransaction, productId, tracker )
    console.log(defaultTransaction.tRowsStack)
    renderTransTable(defaultTransaction)
}


//-----------------------------     PAGE RENDERING   ------------------------//

function initTables(tableId, total)
{
    const tb = document.getElementById(tableId)
    if (tb === null) {
        console.log("unable to find transaction table")
        return 
    }

    tb.innerHTML = `
    <tr class="top-row">
        <th style="width:10%">scan code</th> 
        <th style="width: 40%;">Product Name</th>
        <th style="width: 5%">Qt</th>
        <th style="width: 10%;">Unit Price</th>
        <th style="width: 10%;">discount</th>
        <th style="width: 15%;">Cost</th>
    <th style="width: 5%;">tax</th>
</tr>
`


    for (let i = 0; i < total; i++) {
        let n = (i+1).toString()
        const row = `
        <tr class="trans-row">
                <td id="${'scan-code' + n}" class="trans-row"></td>
                <td id ="${'pname-' + n}" class="trans-row"></td>
                <td id ="${'qt-' + n}" class="trans-row"></td> 
                <td id ="${'unit-price-' + n}" class="trans-row"></td>
                <td id ="${'discount-'+n }" class="trans-row"></td>
                <td id ="${'cost-'+n}" class="trans-row"></td>
                <td id ="${'tax-' + n}" class="trans-row"></td>
            </tr>
        `
        // console.log(row)
        tb.innerHTML += row
    }
}

function renderTransTable(transaction) {
    let tableLen = defalultTableLen
    if (defalultTableLen < defaultTransaction.tRowsStack.length) {
        tableLen = defaultTransaction.tRowsStack.length
    }
    initTables('trans-table', tableLen)

    for (let i = 0; i < transaction.tRowsStack.length; i++) {
        const tr = transaction.tRowsStack[i]
        tr.idx = i+1
        genTableRow(tr, i+1)
        tr.idx = null
    }
    const costDetail = calculateGrandTotal(transaction)
    const subTotalAmount = costDetail.totalCost - costDetail.totalTax

    document.getElementById("total-items").innerText = 
        'Total items: ' + costDetail.totalItems.toString() 
    document.getElementById("sub-total-amount").innerText = 
        'Sub total: ' + subTotalAmount.toString()+ ' 円'
    document.getElementById("grand-total-amount").innerText =　
        'Total: ' + costDetail.totalCost.toString()　+ ' 円'
    document.getElementById("total-tax-amount").innerText = 
        'Total tax: ' +costDetail.totalTax.toString() + ' 円'
}

// rowIndex starts from 1
function genTableRow(tRow, rowIndex) {
    if (tRow.idx === null) {
        console.log("tRow.idx not initiated")
        return
    }
    console.log(tRow.idx)
    const n = rowIndex
    document.getElementById("scan-code" + n).innerText = tRow.pid
    document.getElementById("pname-" + n).innerText = tRow.pName
    document.getElementById("unit-price-" + n).innerText =  tRow.uPrice + ' 円'
    document.getElementById("discount-" + n).innerText = tRow.discount
    document.getElementById("cost-" + n).innerText =    tRow.cost + ' 円'
    document.getElementById("qt-" + n).innerText =  'x ' +  tRow.qty
    document.getElementById("tax-" + n).innerText = tRow.taxAmount
}

function updateTime() {
    // Get the current time
    let currentTime = new Date();
  
    // Get the navbar element
    let navbar = document.getElementById('time-display');
  
    // Set the navbar's innerHTML to the current time
    navbar.innerText = currentTime.toLocaleTimeString();
  }
  
  // Update the time every second
  setInterval(updateTime, 1000);
  

//------------------AJAX request----------------------------------------//
function getProduct() {
    // const scanSound = document.getElementById("scan-audio");
    // scanSound.preload = "auto";
    const input = document.getElementById("in")
    const q = input.value
    var r = new XMLHttpRequest()
    r.open('GET', 'http://localhost:8080/getByID?id='+q)

    r.onload = ()=> {
        
        if (r.status !== 200) {
            console.log('not ok')
            return
        }
       
        const scanButtom = document.getElementById('scan-b')
        scanButtom.disabled = true
        // responseJsonText = r.responseText
        obj = JSON.parse(r.responseText)
        const pd = newProduct(id=obj.ID, pName=obj.Name, price=obj.Price, discount=0, taxP=8)
        addProductToTransaction(defaultTransaction, pd, tracker)
        if (!scanSound.paused) {
            scanSound.currentTime = 0;
            scanSound.play()
        }else{
            scanSound.play()
        }
        scanButtom.disabled = false
        renderTransTable(defaultTransaction)

        return pd
    }
    console.log('sending request')
    r.send()
} 

//------------------ Grand Total---------------------------------

// let grandTotalSpan = document.getElementById('grand-total');
// let subTotalSpan = document.getElementById('sub-total');
// let taxSpan = document.getElementById('tax');

// function updateStatus(grandTotal, subTotal, tax) {
//   grandTotalSpan.innerHTML = 'Grand Total: $' + grandTotal;
//   subTotalSpan.innerHTML = 'Sub Total: $' + subTotal;
//   taxSpan.innerHTML = 'Tax: $' + tax;
// }

// updateStatus(100, 80, 20);




//----------------------------------------------------------------------------------//



const scan = document.getElementById("in");
// scan.addEventListener("keydown", getProduct())

scan.addEventListener("keypress", (event) => {
  if (event.key === "Enter") {
    getProduct()
  }
});

const tracker = new Map()

const defalultTableLen = 20
let tableLen = defalultTableLen
if (defalultTableLen < defaultTransaction.tRowsStack.length) {
    tableLen = defaultTransaction.tRowsStack.length
}
initTables('trans-table', tableLen)
renderTransTable(defaultTransaction)



// Implement KEY PAD

const cashInput = document.getElementById('cash-input')
const productInput = document.getElementById('in')
const productID  = document.getElementById('productID')

var currentFocusedInput = productInput
currentFocusedInput.focus()
for (let i = 0; i < 10; i++) {
    const key = i
    const keypad = document.getElementById('keypad' + key.toString())
    keypad.addEventListener('click', ()=> {
        currentFocusedInput.value += key.toString()
        currentFocusedInput.focus
    })
}
const keypadDelete = document.getElementById('keypad-delete')
keypadDelete.addEventListener('click', ()=> {
    const v = currentFocusedInput.value
    if (v.length > 0) {
        currentFocusedInput.value = v.slice(0, v.length-1)
    }
})


const payForm = document.getElementById('pay-form');
let payButton = document.getElementById('pay-button');

payButton.addEventListener('click', function() {
    currentFocusedInput = document.getElementById('cash-input')
    const detail = calculateGrandTotal(defaultTransaction)
    document.getElementById('payment-due').innerText = 
        'Grand Total: '+ detail.totalCost.toString()
    payForm.style.display = 'block';
    currentFocusedInput.value += '';
})


const completePayment = document.getElementById('complete-payment')
completePayment.addEventListener('click', ()=> {

    let cashRecieved = document.getElementById('cash-input').value
    if (cashRecieved.length > 0) {
        
        const costDetail = calculateGrandTotal(defaultTransaction)
        const cash = parseInt(cashRecieved)
        if (cash > costDetail.totalCost) {
            defaultTransaction.cashRecieved = cash
            defaultTransaction.totalCost = costDetail.totalCost

            //send transaction data
            const transData = JSON.stringify(defaultTransaction)
            console.log(transData)
            
            obj = {
                'pid': null,
                'qty': null,
                'taxP': null,
                'unitPrice': null
            }

        }
    }

})


const cancelPayment = document.getElementById('cancel-payment')
cancelPayment.addEventListener('click', ()=> {
    payForm.style.display = 'none';
    cashInput.value = ''
    currentFocusedInput = productInput
    currentFocusedInput.focus()
})


const registerProduct = document.getElementById('register-product')
const createProductDiv = document.getElementById('create-product')
registerProduct.addEventListener('click', ()=> {
    currentFocusedInput = productID
    createProductDiv.style.display = 'block';
})

const cancelCreateProduct = document.getElementById('cancel-create-product')
cancelCreateProduct.addEventListener('click', ()=>{
    createProductDiv.style.display = 'none';
    currentFocusedInput = productInput
})










// loginForm.addEventListener('click', function(event) {
//   // If the user clicks outside of the login form, hide the login form
//   if (event.target === loginForm) {
//     loginForm.style.display = 'none';
//   }
// });



// console.log(defaultTransaction.tRowsStack)




//-----------------------------      TODO       -------------------------------------------------------//


// TODO:    implement function to add product to transaction 
//          if prdouct exist update  simply update tRow which has the product

// TODO:    Implement input of product to transaction (Barcode scan would have been perfect but)
//          create input bar to input product-id  and pressing enter or clicking input 
//          should imput the product-id. If product had not found id database. Notify user (sound or blink)
//                   

//TODO:     Delte t-row from transaction // Delte the entire t-row /reset

// TODO:    Create a div to show current grand total total tax total discount 

// TODO:    Implement a function to complete transaction ( Cash Recived total Change)
//          create a button pay-by-cash which when clicked Popup new small window which has
//          input of cash amount and has button to complete transaction


// TODO:   Before transaction  can be completed send transaction data to backend to verify the 
//         transaction integrity 

//