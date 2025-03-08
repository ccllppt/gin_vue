<template>
    <div class="LoginView">
        <b-row class="mt-5">
        <b-col
        md="8"
        offset-md="3"
        lg="6"
        offset-lg="3">
        <b-card title="Login">

        <b-form>
      <b-form-group label="Telephone">
        <b-form-input
          v-model="$v.user.telephone.$model"
          type="number"
          placeholder="Telephone"
          :state= "validateState('telephone')"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState('telephone')">
          The phone number does not meet the requirements
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="Password">
        <b-form-input
          v-model="$v.user.password.$model"
          type="password"
          placeholder="Password"
          :state= "validateState('password')"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState('password')">
          The password does not meet the requirements
        </b-form-invalid-feedback>
        <b-button @click="login" class="mt-4"
         variant="outline-primary" block>Login</b-button>
      </b-form-group>
      </b-form>
    </b-card>
    </b-col>
    </b-row>
    </div>
</template>

<script lang="ts">
import {
  required, minLength,
} from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';
import { mapMutations } from 'vuex';

export default ({
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator.telephonrValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },

    },

  },
  methods: {
    ...mapMutations('userModule', ['SET_TOKEN', 'SET_USERINFO']),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.$store.dispatch('userModule/login', this.user).then(() => {
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
});

</script>

<style scoped>

</style>
