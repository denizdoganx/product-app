window.ProductApp = {};

ProductApp.ApiTestModule = () => {
    const self = {};

    self.init = async () => {
        self.setEvents();
    };

    self.setEvents = () => {
        document.getElementById('get-products-button').addEventListener('click', async () => {
            const response = await fetch("http://localhost:9000/get-all-products");
            const data = await response.json();

            const list = document.getElementById("product-list");
            list.innerHTML = "";

            data.forEach(({ name, price }) => {
                const li = document.createElement("li");
                li.innerText = `${ name } - ${ price }`;
                list.appnedChild(li);
            });
        });
    };

    self.init();
};

ProductApp.ApiTestModule();