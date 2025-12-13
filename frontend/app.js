ProductApp.ApiTestModule = () => {
    const self = {};

    self.init = async () => {
        self.setEvents();
    };

    self.setEvents = () => {
        document.getElementById('get-products-button').addEventListener('click', async () => {
            const response = await fetch("http://localhost:9000/api/v1/products");
            const htmlText = await response.text();

            const list = document.getElementById("product-list");
            list.innerText = htmlText;
        });
    };

    self.init();
};

ProductApp.ApiTestModule();