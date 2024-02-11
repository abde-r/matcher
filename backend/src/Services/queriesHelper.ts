const pool = require('../../database/dbConfig');

const CREATE = async (args: any) => {
    const { table, columns } = args;
    let query: string = `CREATE TABLE IF NOT EXISTS "${table}" ( id SERIAL PRIMARY KEY, `;
    let i: number=0;
    
    columns.forEach((column: any) => {
        query+=`${column.name} ${column.type}  ${column.default}`;
        if (i+1 < columns.length)
            query+=`, `;
        i++;
    });
    query+=`);`;
    try {
        const client = await pool.connect();
        await client.query(query);
        console.log(`✅-> Table ${table} Created Successfully!`);
    }
    catch (err) {
        console.error('❌-> Error Creating Table', table, err);
    }
}

const INSERT = async (args: any) => {
    const { table, columns } = args;
    let query: string = `INSERT INTO "${table}" (`;
    let i: number=0;
    const values: any = [];
    
    columns.forEach((column: any) => {
        query+=`${column.name}`;
        if (i+1 < columns.length)
            query+=`, `;
        i++;
    });
    query+=`) VALUES (`;
    i=0;
    columns.forEach((column: any) => {
        query+=`$${i+1}`;
        values.push(column.value);
        if (i+1 < columns.length)
            query+=`, `;
        i++;
    });
    query+=`);`;
    console.log('ha lquery: ', query)
    try {
        const client = await pool.connect();
        await client.query(query, values);
    }
    catch (err) {
        console.error('❌-> Error Accessing Table', table, err);
    }
}

const UPDATE = async (args: any) => {

}

const SELECT = async (args: any) => {
    const { table, columns } = args;
    let query: string = `SELECT `;
    
    if (columns.length) {
        columns.forEach((column: any) => {
            query+=`${column} `;
        });
    }
    else
        query+='* ';
    query+=`FROM "${table}";`;
    try {
        const client = await pool.connect();
        return await client.query(query);
    }
    catch (err) {
        console.error('❌-> Error Accessing Table', table, err);
    }    
}

const DELETE = async (args: any) => {
    
}

module.exports = { CREATE, SELECT, INSERT };